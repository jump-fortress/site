package internal

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"regexp"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

var (
	steamID64Offset = uint64(76561197960265728)

	displayNameRegex = regexp.MustCompile(`^(([[:word:]]|[[:punct:]])+[[:space:]]?)*([[:word:]]|[[:punct:]])+$`)
	urlRegex         = regexp.MustCompile(`^https:\/\/steamcommunity.com\/tradeoffer\/new\/\?partner=([[:digit:]])+&token=([[:word:]]|\-)+$`)
)

func getSteamIDFromSteamID64(steamID64 uint64) string {
	// Z = steam_id64 - 76561197960265728
	// Z = (Z - 1) / 2
	// STEAM_0:1:Z
	steamID_Z := (steamID64 - steamID64Offset - 1) / 2
	steamID_Z_string := strconv.FormatUint(steamID_Z, 10)
	steamID := fmt.Sprintf("STEAM_0:1:%s", steamID_Z_string)

	return steamID
}

func getTempusPlayerInfo(tempusID int64) (*responses.TempusPlayerInfo, error) {
	url := fmt.Sprintf("https://tempus2.xyz/api/v0/players/id/%d/info", tempusID)

	response, err := retryablehttp.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	tempusPlayer := &responses.TempusPlayerInfo{}
	if err := json.Unmarshal(body, &tempusPlayer); err != nil {
		return nil, err
	}

	return tempusPlayer, nil
}

// todo: check some nullable before using them?
func playerResponseFromPlayer(player queries.Player) responses.Player {
	return responses.Player{
		ID:                player.ID,
		Role:              player.Role,
		SteamAvatarUrl:    player.SteamAvatarUrl.String,
		TempusID:          player.TempusID.Int64,
		DisplayName:       player.DisplayName.String,
		SoldierDivision:   player.SoldierDivision.String,
		DemoDivision:      player.DemoDivision.String,
		PreferredClass:    player.PreferredClass,
		PreferredLauncher: player.PreferredLauncher,
		CreatedAt:         player.CreatedAt,
	}
}

func fullPlayerResponseFromPlayer(player queries.Player) responses.FullPlayer {
	return responses.FullPlayer{
		ID:                player.ID,
		Role:              player.Role,
		SteamId64:         player.SteamId64,
		SteamAvatarUrl:    player.SteamAvatarUrl.String,
		SteamTradeToken:   player.SteamTradeToken.String,
		TempusID:          player.TempusID.Int64,
		DiscordID:         player.DiscordID.String,
		DisplayName:       player.DisplayName.String,
		SoldierDivision:   player.SoldierDivision.String,
		DemoDivision:      player.DemoDivision.String,
		PreferredClass:    player.PreferredClass,
		PreferredLauncher: player.PreferredLauncher,
		CreatedAt:         player.CreatedAt,
	}
}

func HandleGetPlayer(ctx context.Context, _ *struct{}) (*responses.FullPlayerOutput, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	steamID64_string := strconv.FormatUint(principal.SteamID, 10)

	player, err := responses.Queries.SelectPlayerFromSteamID64(ctx, steamID64_string)
	if err != nil {
		return nil, err
	}

	resp := &responses.FullPlayerOutput{Body: fullPlayerResponseFromPlayer(player)}
	return resp, nil
}

func HandleGetPlayerByID(ctx context.Context, input *responses.PlayerIDInput) (*responses.PlayerOutput, error) {
	player, err := responses.Queries.SelectPlayer(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayerOutput{Body: playerResponseFromPlayer(player)}
	return resp, nil
}

func HandleGetPlayerBySteamID64(ctx context.Context, input *responses.PlayerSteamID64Input) (*responses.PlayerOutput, error) {
	steamID64_string := strconv.FormatUint(input.SteamID64, 10)

	player, err := responses.Queries.SelectPlayerFromSteamID64(ctx, steamID64_string)
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayerOutput{Body: playerResponseFromPlayer(player)}
	return resp, nil
}

func HandleGetPlayerProfileByID(ctx context.Context, input *responses.PlayerIDInput) (*responses.PlayerProfileOutput, error) {

	player, err := responses.Queries.SelectPlayer(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	soldierPoints, err := responses.Queries.SelectPlayerPoints(ctx, queries.SelectPlayerPointsParams{
		PlayerID: player.ID,
		Class:    "Soldier",
	})
	if err != nil {
		return nil, err
	}

	demoPoints, err := responses.Queries.SelectPlayerPoints(ctx, queries.SelectPlayerPointsParams{
		PlayerID: player.ID,
		Class:    "Demo",
	})
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayerProfileOutput{
		Body: responses.PlayerProfile{
			Player: playerResponseFromPlayer(player),
			SoldierPoints: responses.PlayerPoints{
				Total:        soldierPoints.Total,
				Last3Monthly: soldierPoints.Last3Monthly,
				Last9Motw:    soldierPoints.Last9Motw,
			},
			DemoPoints: responses.PlayerPoints{
				Total:        demoPoints.Total,
				Last3Monthly: demoPoints.Last3Monthly,
				Last9Motw:    demoPoints.Last9Motw,
			},
		},
	}
	return resp, nil
}

func HandlePutPlayerPreferredClass(ctx context.Context, input *responses.ClassNameInput) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}
	steamID64_string := strconv.FormatUint(principal.SteamID, 10)
	if err := responses.Queries.UpdatePlayerPreferredClassFromSteamID64(ctx, queries.UpdatePlayerPreferredClassFromSteamID64Params{
		PreferredClass: input.Class,
		SteamId64:      steamID64_string,
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func HandlePutPlayerPreferredLauncher(ctx context.Context, input *responses.LauncherNameInput) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}
	steamID64_string := strconv.FormatUint(principal.SteamID, 10)
	if err := responses.Queries.UpdatePlayerPreferredLauncherFromSteamID64(ctx, queries.UpdatePlayerPreferredLauncherFromSteamID64Params{
		PreferredLauncher: input.Launcher,
		SteamId64:         steamID64_string,
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func HandleGetAllPlayers(ctx context.Context, _ *struct{}) (*responses.ManyPlayersOutput, error) {
	players, err := responses.Queries.SelectAllPlayers(ctx)
	if err != nil {
		return nil, err
	}

	resp := &responses.ManyPlayersOutput{
		Body: []responses.Player{},
	}

	for _, p := range players {
		playerResponse := playerResponseFromPlayer(p)
		resp.Body = append(resp.Body, playerResponse)
	}
	return resp, nil
}

func HandleGetAllPlayersFull(ctx context.Context, _ *struct{}) (*responses.ManyFullPlayersOutput, error) {
	players, err := responses.Queries.SelectAllPlayers(ctx)
	if err != nil {
		return nil, err
	}

	resp := &responses.ManyFullPlayersOutput{
		Body: []responses.FullPlayer{},
	}

	for _, p := range players {
		fullPlayerResponse := fullPlayerResponseFromPlayer(p)
		resp.Body = append(resp.Body, fullPlayerResponse)
	}

	return resp, nil
}

func HandlePutPlayerDisplayName(ctx context.Context, input *responses.DisplayNameInput) (*struct{}, error) {
	// todo: check for string validity
	if len(input.Name) > 32 {
		return nil, huma.Error400BadRequest("display name is too long (max 32 characters)")
	}
	if !displayNameRegex.MatchString(input.Name) {
		return nil, huma.Error400BadRequest("display name is not in the expected format. please use alphanumeric, spaces, and punctuation only.")
	}

	if err := responses.Queries.UpdatePlayerDisplayName(ctx, queries.UpdatePlayerDisplayNameParams{
		DisplayName: sql.NullString{
			String: input.Name,
			Valid:  true,
		},
		ID: input.ID,
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func HandlePutPlayerSteamTradeToken(ctx context.Context, input *responses.SteamTradeURL) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}
	steamID64_string := strconv.FormatUint(principal.SteamID, 10)

	// todo: check for string validity
	if len(input.Url) > 100 {
		return nil, huma.Error400BadRequest("URL is too long (max 100 characters)")
	}

	if !urlRegex.MatchString(input.Url) {
		return nil, huma.Error400BadRequest("URL is not in the expected format. Please paste the full Steam Trade URL.")
	}

	parsedInputUrl, err := url.Parse(input.Url)
	if err != nil {
		return nil, huma.Error400BadRequest("URL couldn't be resolved. Please paste the full Steam Trade URL.")
	}

	if parsedInputUrl.Host != "steamcommunity.com" || parsedInputUrl.Path != "/tradeoffer/new/" {
		return nil, huma.Error400BadRequest("URL couldn't be resolved. Please paste the full Steam Trade URL.")
	}

	params := parsedInputUrl.Query()
	parsedSteamID3 := params.Get("partner")
	parsedSteamTradeToken := params.Get("token")

	if parsedSteamID3 == "" || parsedSteamTradeToken == "" {
		return nil, huma.Error400BadRequest("URL couldn't be resolved. Please paste the full Steam Trade URL.")
	}

	steamID3 := strconv.FormatUint(principal.SteamID-steamID64Offset, 10)
	if steamID3 != parsedSteamID3 {
		return nil, huma.Error400BadRequest("This URL didn't match your SteamID3. Please check if you're logged in to the correct Steam account.")
	}

	if err := responses.Queries.UpdatePlayerSteamTradeTokenFromSteamID64(ctx, queries.UpdatePlayerSteamTradeTokenFromSteamID64Params{
		SteamTradeToken: sql.NullString{
			String: parsedSteamTradeToken,
			Valid:  true,
		},
		SteamId64: steamID64_string,
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func HandlePutPlayerTempusID(ctx context.Context, input *responses.TempusIDInput) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	steamID64_string := strconv.FormatUint(principal.SteamID, 10)

	player, err := responses.Queries.SelectPlayerFromSteamID64(ctx, steamID64_string)
	if err != nil {
		return nil, huma.Error500InternalServerError("Something went wrong while checking if your Tempus ID was already set..")
	}

	if player.TempusID.Valid {
		return nil, huma.Error400BadRequest(fmt.Sprintf("Your Tempus ID is already set to %d", player.TempusID.Int64))
	}

	steamID := getSteamIDFromSteamID64(principal.SteamID)

	tempusPlayer, err := getTempusPlayerInfo(input.TempusID)
	if err != nil {
		return nil, huma.Error503ServiceUnavailable("Couldn't verify this Tempus ID with Tempus. If Tempus isn't down, please check your Tempus ID again.")
	}

	if steamID != tempusPlayer.SteamID {
		return nil, huma.Error400BadRequest(fmt.Sprintf("This player's (%s) Steam ID doesn't match the Tempus ID you provided.", tempusPlayer.TempusName))
	}

	if err := responses.Queries.UpdatePlayerTempusIDFromSteamID64(ctx, queries.UpdatePlayerTempusIDFromSteamID64Params{
		TempusID: sql.NullInt64{
			Int64: input.TempusID,
			Valid: true,
		},
		SteamId64: steamID64_string,
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func HandlePutPlayerSteamAvatarUrl(ctx context.Context, _ *struct{}) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	steamProfile, err := FetchProfileSummary(principal.SteamID)
	if err != nil {
		return nil, huma.Error503ServiceUnavailable("Couldn't get your profile from Steam. If Steam isn't down, please try again.")
	}

	if err = responses.Queries.UpdatePlayerSteamAvatarURLFromSteamID64(ctx, queries.UpdatePlayerSteamAvatarURLFromSteamID64Params{
		SteamAvatarUrl: sql.NullString{
			String: steamProfile.AvatarFullURL,
			Valid:  true,
		},
		SteamId64: steamProfile.SteamID,
	}); err != nil {
		return nil, err
	}

	return nil, nil
}
