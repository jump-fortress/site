package internal

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/rotisserie/eris"
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

var (
	displayNameRegex = regexp.MustCompile(`^(([[:word:]]|[[:punct:]])+[[:space:]]?)*([[:word:]]|[[:punct:]])+$`)
)

func getTempusPlayerInfo(tempusID int64) (*responses.TempusPlayerInfo, error) {
	url := fmt.Sprintf("https://tempus2.xyz/api/v0/players/id/%d/stats", tempusID)

	response, err := retryablehttp.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	tempusResponsePlayerInfo := &responses.TempusResponsePlayerInfo{}
	if err := json.Unmarshal(body, &tempusResponsePlayerInfo); err != nil {
		return nil, err
	}

	return &tempusResponsePlayerInfo.PlayerInfo, nil
}

// todo: check some nullable before using them?
func playerResponseFromPlayer(player queries.Player) responses.Player {
	return responses.Player{
		ID:                player.ID,
		Role:              player.Role,
		SteamAvatarUrl:    player.SteamAvatarUrl.String,
		TempusID:          player.TempusID.Int64,
		Country:           player.Country.String,
		CountryCode:       player.CountryCode.String,
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
		SteamAvatarUrl:    player.SteamAvatarUrl.String,
		SteamTradeToken:   player.SteamTradeToken.String,
		TempusID:          player.TempusID.Int64,
		Country:           player.Country.String,
		CountryCode:       player.CountryCode.String,
		DiscordID:         player.DiscordID.String,
		DisplayName:       player.DisplayName.String,
		SoldierDivision:   player.SoldierDivision.String,
		DemoDivision:      player.DemoDivision.String,
		PreferredClass:    player.PreferredClass,
		PreferredLauncher: player.PreferredLauncher,
		CreatedAt:         player.CreatedAt,
	}
}

func HandleGetSelfPlayer(ctx context.Context, _ *struct{}) (*responses.FullPlayerOutput, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	player, err := responses.Queries.SelectPlayer(ctx, principal.SteamID.String())
	if err != nil {
		return nil, err
	}

	resp := &responses.FullPlayerOutput{Body: fullPlayerResponseFromPlayer(player)}
	return resp, nil
}

func HandleGetPlayer(ctx context.Context, input *responses.PlayerIDInput) (*responses.PlayerOutput, error) {
	player, err := responses.Queries.SelectPlayer(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayerOutput{Body: playerResponseFromPlayer(player)}
	return resp, nil
}

func HandleGetPlayerProfile(ctx context.Context, input *responses.PlayerIDInput) (*responses.PlayerProfileOutput, error) {

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

func HandlePutSelfPreferredClass(ctx context.Context, input *responses.ClassNameInput) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	if err := responses.Queries.UpdatePlayerPreferredClass(ctx, queries.UpdatePlayerPreferredClassParams{
		PreferredClass: input.Class,
		ID:             principal.SteamID.String(),
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func HandlePutSelfPreferredLauncher(ctx context.Context, input *responses.LauncherNameInput) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	if err := responses.Queries.UpdatePlayerPreferredLauncher(ctx, queries.UpdatePlayerPreferredLauncherParams{
		PreferredLauncher: input.Launcher,
		ID:                principal.SteamID.String(),
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

func HandleGetAllFullPlayers(ctx context.Context, _ *struct{}) (*responses.ManyFullPlayersOutput, error) {
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

func HandlePutSelfSteamTradeToken(ctx context.Context, input *responses.SteamTradeURL) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	if len(input.Url) > 100 {
		return nil, huma.Error400BadRequest("URL is too long")
	}

	// extract steam3ID and token from input
	var scannedSteam3ID uint32
	var token string
	_, err := fmt.Sscanf(input.Url, "https://steamcommunity.com/tradeoffer/new/?partner=%d&token=%s", &scannedSteam3ID, &token)
	if err != nil {
		return nil, huma.Error400BadRequest("URL couldn't be resolved. Please paste the full Steam Trade URL.")
	}

	if principal.SteamID.AccountId() != scannedSteam3ID {
		return nil, huma.Error400BadRequest("This URL didn't match your SteamID. Please check if you're logged in to the correct Steam account.")
	}

	if err := responses.Queries.UpdatePlayerSteamTradeToken(ctx, queries.UpdatePlayerSteamTradeTokenParams{
		SteamTradeToken: sql.NullString{
			String: token,
			Valid:  true,
		},
		ID: principal.SteamID.String(),
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func HandlePutSelfTempusInfo(ctx context.Context, input *responses.TempusIDInput) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	player, err := responses.Queries.SelectPlayer(ctx, principal.SteamID.String())
	if err != nil {
		return nil, huma.Error500InternalServerError("Something went wrong while checking if your Tempus ID was already set..")
	}

	if player.TempusID.Valid {
		return nil, huma.Error400BadRequest("Tempus ID already set")
	}

	tempusPlayer, err := getTempusPlayerInfo(input.TempusID)
	if err != nil {
		return nil, huma.Error503ServiceUnavailable("Couldn't verify this Tempus ID with Tempus. If Tempus isn't down, please check your Tempus ID again.")
	}

	var tempusPlayerSteamID3 int
	_, err = fmt.Sscanf(tempusPlayer.SteamID, "STEAM_0:1:%d", &tempusPlayerSteamID3)
	if err != nil {
		return nil, huma.Error503ServiceUnavailable("Couldn't verify this Tempus ID with Tempus due to Tempus's response. Is Tempus functioning correctly?")
	}

	tempusPlayerSteamID3 = tempusPlayerSteamID3*2 + 1
	if int(principal.SteamID.AccountId()) != tempusPlayerSteamID3 {
		return nil, huma.Error400BadRequest(fmt.Sprintf("This player's (%s) Steam ID doesn't match the Tempus ID you provided.", tempusPlayer.TempusName))
	}

	if err := responses.Queries.UpdatePlayerTempusInfo(ctx, queries.UpdatePlayerTempusInfoParams{
		TempusID: sql.NullInt64{
			Int64: input.TempusID,
			Valid: true,
		},
		Country: sql.NullString{
			String: tempusPlayer.Country,
			Valid:  true,
		},
		CountryCode: sql.NullString{
			String: strings.ToLower(tempusPlayer.CountryCode),
			Valid:  true,
		},
		ID: principal.SteamID.String(),
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func HandlePutSelfSteamAvatarUrl(ctx context.Context, _ *struct{}) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	steamProfile, err := FetchProfileSummary(principal.SteamID.ID())
	if err != nil {
		return nil, huma.Error503ServiceUnavailable("Couldn't get your profile from Steam. If Steam isn't down, please try again.")
	}

	if err = responses.Queries.UpdatePlayerSteamAvatarURL(ctx, queries.UpdatePlayerSteamAvatarURLParams{
		SteamAvatarUrl: sql.NullString{
			String: steamProfile.AvatarFullURL,
			Valid:  true,
		},
		ID: principal.SteamID.String(),
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

// todo: implement
func HandlePutSelfPlayerRequest(ctx context.Context, input *responses.PlayerRequestInput) (*struct{}, error) {
	// check for an existing request..

	// if display name change
	// do it

	// if division placement
	// check for linked tempus id, and make sure player doesn't have the division they're requesting
	// do it
	return nil, eris.New("todo")
}
