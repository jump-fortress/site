package internal

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

var (
	// alphanumeric, allow non-leading, non-repeating underscores, whitespace, and dashes
	displayNameRegex = regexp.MustCompile(`^((([a-z]|[A-Z]|\d|\.)+(_|\ |\-)?)+)*([a-z]|[A-Z]|\d|\.)+$`)
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

	tempusResponsePlayerInfo := &responses.TempusPlayerInfoResponse{}
	if err := json.Unmarshal(body, &tempusResponsePlayerInfo); err != nil {
		return nil, err
	}

	return &tempusResponsePlayerInfo.PlayerInfo, nil
}

// todo: check some nullable before using them?
func getPlayerPreviewResponse(player queries.Player) responses.PlayerPreview {
	return responses.PlayerPreview{
		ID:                player.ID,
		Role:              player.Role,
		SteamAvatarUrl:    player.SteamAvatarUrl.String,
		TempusID:          player.TempusID.Int64,
		Country:           player.Country.String,
		CountryCode:       player.CountryCode.String,
		DisplayName:       player.DisplayName.String,
		SoldierDivision:   player.SoldierDivision.String,
		DemoDivision:      player.DemoDivision.String,
		MotwTimeslot:      player.MotwTimeslot.Int64,
		PreferredClass:    player.PreferredClass,
		PreferredLauncher: player.PreferredLauncher.String,
		PreferredMap:      player.PreferredMap.String,
		CreatedAt:         player.CreatedAt,
	}
}

func getPlayerResponse(player queries.Player) responses.Player {
	return responses.Player{
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
		MotwTimeslot:      player.MotwTimeslot.Int64,
		PreferredClass:    player.PreferredClass,
		PreferredLauncher: player.PreferredLauncher.String,
		PreferredMap:      player.PreferredMap.String,
		CreatedAt:         player.CreatedAt,
	}
}

func getPlayerRequestPreviewResponse(request queries.PlayerRequest) responses.PlayerRequestPreview {
	return responses.PlayerRequestPreview{
		RequestType:   request.Type,
		RequestString: request.Content.String,
		CreatedAt:     request.CreatedAt,
	}
}

func getPlayerWithRequestResponse(request queries.SelectAllPendingPlayerRequestsRow) responses.PlayerWithRequest {
	return responses.PlayerWithRequest{
		Request: responses.PlayerRequest{
			ID:            request.ID,
			PlayerID:      request.PlayerID,
			RequestType:   request.Type,
			RequestString: request.Content.String,
			Pending:       request.Pending,
			CreatedAt:     request.CreatedAt,
		},
		Player: responses.Player{
			ID:                request.ID_2,
			Role:              request.Role,
			SteamAvatarUrl:    request.SteamAvatarUrl.String,
			SteamTradeToken:   request.SteamTradeToken.String,
			TempusID:          request.TempusID.Int64,
			Country:           request.Country.String,
			CountryCode:       request.CountryCode.String,
			DiscordID:         request.DiscordID.String,
			DisplayName:       request.DisplayName.String,
			SoldierDivision:   request.SoldierDivision.String,
			DemoDivision:      request.DemoDivision.String,
			PreferredClass:    request.PreferredClass,
			PreferredLauncher: request.PreferredLauncher.String,
			CreatedAt:         request.CreatedAt_2,
		},
	}
}

func HandleGetSelfPlayer(ctx context.Context, _ *struct{}) (*responses.PlayerOutput, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	player, err := responses.Queries.SelectPlayer(ctx, principal.SteamID.String())
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayerOutput{Body: getPlayerResponse(player)}
	return resp, nil
}

func HandleGetPlayer(ctx context.Context, input *responses.PlayerIDInput) (*responses.PlayerPreviewOutput, error) {
	player, err := responses.Queries.SelectPlayer(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayerPreviewOutput{Body: getPlayerPreviewResponse(player)}
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
			PlayerPreview: getPlayerPreviewResponse(player),
			PlayerPoints: responses.PlayerPoints{
				SoldierPoints: responses.PlayerClassPoints{
					Total:        soldierPoints.Total,
					Last3Monthly: soldierPoints.Last3Monthly,
					Last9Motw:    soldierPoints.Last9Motw,
				},
				DemoPoints: responses.PlayerClassPoints{
					Total:        demoPoints.Total,
					Last3Monthly: demoPoints.Last3Monthly,
					Last9Motw:    demoPoints.Last9Motw,
				},
			},
		},
	}
	return resp, nil
}

func HandlePostSelfPreferredClass(ctx context.Context, input *responses.ClassNameInput) (*struct{}, error) {
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

func HandlePostSelfPreferredLauncher(ctx context.Context, input *responses.LauncherNameInput) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	noLauncher := input.Launcher == "None"

	if err := responses.Queries.UpdatePlayerPreferredLauncher(ctx, queries.UpdatePlayerPreferredLauncherParams{
		PreferredLauncher: sql.NullString{
			String: input.Launcher,
			Valid:  !noLauncher,
		},
		ID: principal.SteamID.String(),
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func HandlePostSelfPreferredMap(ctx context.Context, input *responses.MapNameInput) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	noMap := input.Map == "none"

	maps, err := responses.Queries.SelectMapNames(ctx)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(maps, input.Map) && !noMap {
		return nil, huma.Error400BadRequest("invalid map name")
	}

	if err := responses.Queries.UpdatePlayerPreferredMap(ctx, queries.UpdatePlayerPreferredMapParams{
		PreferredMap: sql.NullString{
			String: input.Map,
			Valid:  !noMap,
		},
		ID: principal.SteamID.String(),
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func HandleGetAllPlayers(ctx context.Context, _ *struct{}) (*responses.PlayerPreviewsOutput, error) {
	players, err := responses.Queries.SelectAllPlayers(ctx)
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayerPreviewsOutput{
		Body: []responses.PlayerPreview{},
	}

	for i, p := range players {

		if p.Country.String == "" {
			fmt.Println(strconv.Itoa(i) + ": attempting for " + strconv.Itoa(int(p.TempusID.Int64)))
			t, err := getTempusPlayerInfo(p.TempusID.Int64)
			time.Sleep(1 * time.Second)
			if err != nil {
				continue
			}
			err = responses.Queries.UpdatePlayerTempusInfo(ctx, queries.UpdatePlayerTempusInfoParams{
				TempusID: sql.NullInt64{
					Int64: p.TempusID.Int64,
					Valid: true,
				},
				Country: sql.NullString{
					String: t.Country,
					Valid:  true,
				},
				CountryCode: sql.NullString{
					String: strings.ToLower(t.CountryCode),
					Valid:  true,
				},
				ID: p.ID,
			})
			if err != nil {
				continue
			}
		}

		playerResponse := getPlayerPreviewResponse(p)
		resp.Body = append(resp.Body, playerResponse)
	}

	return resp, nil
}

func HandlePostSelfSteamTradeToken(ctx context.Context, input *responses.SteamTradeURL) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	if len(input.Url) > 100 {
		return nil, huma.Error400BadRequest("URL is too long. please paste the full Steam Trade URL")
	}

	// extract steam3ID and token from input
	var scannedSteam3ID uint32
	var token string
	_, err := fmt.Sscanf(input.Url, "https://steamcommunity.com/tradeoffer/new/?partner=%d&token=%s", &scannedSteam3ID, &token)
	if err != nil {
		return nil, huma.Error400BadRequest("URL couldn't be resolved. please paste the full Steam Trade URL.")
	}

	if principal.SteamID.AccountId() != scannedSteam3ID {
		return nil, huma.Error400BadRequest(fmt.Sprintf("URL didn't match your profile's SteamID (%d). please check if you're logged in to the correct Steam account.", scannedSteam3ID))
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

// todo: notify to set this player's divisions when they link their tempus id
func HandlePostSelfTempusInfo(ctx context.Context, input *responses.TempusIDInput) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	player, err := responses.Queries.SelectPlayer(ctx, principal.SteamID.String())
	if err != nil {
		return nil, huma.Error500InternalServerError("something went wrong when checking if your Tempus ID was set..")
	}

	if player.TempusID.Valid {
		return nil, huma.Error400BadRequest("already set")
	}

	tempusPlayer, err := getTempusPlayerInfo(input.TempusID)
	if err != nil {
		return nil, huma.Error503ServiceUnavailable("couldn't verify with Tempus. if Tempus isn't down, please check your Tempus ID again")
	}

	var tempusPlayerSteamID3 int
	_, err = fmt.Sscanf(tempusPlayer.SteamID, "STEAM_0:1:%d", &tempusPlayerSteamID3)
	if err != nil {
		return nil, huma.Error503ServiceUnavailable("couldn't verify with Tempus due to Tempus's response. is Tempus functioning correctly?")
	}

	tempusPlayerSteamID3 = tempusPlayerSteamID3*2 + 1
	if int(principal.SteamID.AccountId()) != tempusPlayerSteamID3 {
		return nil, huma.Error400BadRequest(fmt.Sprintf("Tempus player (%s) SteamID doesn't match your SteamID", tempusPlayer.TempusName))
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

func HandlePostSelfSteamAvatarUrl(ctx context.Context, _ *struct{}) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	steamProfile, err := FetchProfileSummary(principal.SteamID.ID())
	if err != nil {
		return nil, huma.Error503ServiceUnavailable("couldn't get your profile from Steam. If Steam isn't down, please try again.")
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

// player requests..
// 1. make sure there are no pending requests for this request type
// 2. create the request if it's for a display name change
// 3. make sure Tempus ID is linked and the division requested is missing
func HandlePostSelfPlayerRequest(ctx context.Context, input *responses.PlayerRequestInput) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	// check for an existing request..
	requestExists, err := responses.Queries.CheckPendingPlayerRequest(ctx, queries.CheckPendingPlayerRequestParams{
		PlayerID: principal.SteamID.String(),
		Type:     input.RequestType,
	})
	if err != nil {
		return nil, err
	}

	if requestExists == 1 {
		return nil, huma.Error409Conflict(fmt.Sprintf("a %s request already exists", input.RequestType))
	} else {

		if input.RequestType == "Display Name Change" {
			if err := responses.Queries.InsertPlayerRequest(ctx, queries.InsertPlayerRequestParams{
				PlayerID: principal.SteamID.String(),
				Type:     input.RequestType,
				Content: sql.NullString{
					String: input.RequestString,
					Valid:  true,
				},
			}); err != nil {
				return nil, huma.Error500InternalServerError("something went wrong creating this request")
			}

			// division placement request
		} else {
			player, err := responses.Queries.SelectPlayer(ctx, principal.SteamID.String())
			if err != nil {
				return nil, huma.Error500InternalServerError("something went wrong creating this request")
			}

			if !player.TempusID.Valid {
				return nil, huma.Error400BadRequest("please link your Tempus ID first")
			}

			if input.RequestType == "Soldier Placement" {
				if player.SoldierDivision.Valid {
					return nil, huma.Error400BadRequest("you already have a soldier division")
				}
				// demo placement
			} else {
				if player.DemoDivision.Valid {
					return nil, huma.Error400BadRequest("you already have a demo division")
				}
			}

			if err := responses.Queries.InsertPlayerRequest(ctx, queries.InsertPlayerRequestParams{
				PlayerID: principal.SteamID.String(),
				Type:     input.RequestType,
				Content: sql.NullString{
					String: "",
					Valid:  false,
				},
			}); err != nil {
				return nil, huma.Error500InternalServerError("something went wrong creating this request")
			}
		}
	}

	return nil, nil
}

func HandleGetSelfPlayerRequests(ctx context.Context, _ *struct{}) (*responses.PlayerRequestPreviewsOutput, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	requests, err := responses.Queries.SelectPendingPlayerRequests(ctx, principal.SteamID.String())
	if err != nil {
		return nil, nil
	}

	resp := &responses.PlayerRequestPreviewsOutput{
		Body: []responses.PlayerRequestPreview{},
	}

	for _, r := range requests {
		fullRequestResponse := getPlayerRequestPreviewResponse(r)
		resp.Body = append(resp.Body, fullRequestResponse)
	}
	return resp, nil
}

// consultant

func HandleGetAllFullPlayers(ctx context.Context, _ *struct{}) (*responses.PlayersOutput, error) {
	players, err := responses.Queries.SelectAllPlayers(ctx)
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayersOutput{
		Body: []responses.Player{},
	}

	for _, p := range players {
		fullPlayerResponse := getPlayerResponse(p)
		resp.Body = append(resp.Body, fullPlayerResponse)
	}

	return resp, nil
}

func HandleGetAllPendingPlayerRequests(ctx context.Context, _ *struct{}) (*responses.PlayersWithRequestOutput, error) {
	requests, err := responses.Queries.SelectAllPendingPlayerRequests(ctx)
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayersWithRequestOutput{
		Body: []responses.PlayerWithRequest{},
	}

	for _, r := range requests {
		requestResponse := getPlayerWithRequestResponse(r)
		resp.Body = append(resp.Body, requestResponse)
	}

	return resp, nil
}

// moderator

func HandlePostPlayerDisplayName(ctx context.Context, input *responses.DisplayNameInput) (*struct{}, error) {
	// todo: check for string validity
	if len(input.Name) > 32 {
		return nil, huma.Error400BadRequest("too long (max 32 characters)")
	}
	if !displayNameRegex.MatchString(input.Name) {
		return nil, huma.Error400BadRequest("please use only alphanumeric, spaces, and punctuation characters.")
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

func HandlePostPlayerSoldierDivision(ctx context.Context, input *responses.UpdateDivisionInput) (*struct{}, error) {
	if input.Division == "None" {
		if err := responses.Queries.UpdatePlayerSoldierDivision(ctx, queries.UpdatePlayerSoldierDivisionParams{
			SoldierDivision: sql.NullString{
				String: "",
				Valid:  false,
			},
			ID: input.ID,
		}); err != nil {
			return nil, err
		}

	} else {
		if !slices.Contains(responses.Divisions, input.Division) {
			return nil, huma.Error400BadRequest(fmt.Sprintf("%s isn't a valid division.", input.Division))
		}

		if err := responses.Queries.UpdatePlayerSoldierDivision(ctx, queries.UpdatePlayerSoldierDivisionParams{
			SoldierDivision: sql.NullString{
				String: input.Division,
				Valid:  true,
			},
			ID: input.ID,
		}); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func HandlePostPlayerDemoDivision(ctx context.Context, input *responses.UpdateDivisionInput) (*struct{}, error) {
	if input.Division == "None" {
		if err := responses.Queries.UpdatePlayerDemoDivision(ctx, queries.UpdatePlayerDemoDivisionParams{
			DemoDivision: sql.NullString{
				String: "",
				Valid:  false,
			},
			ID: input.ID,
		}); err != nil {
			return nil, err
		}

	} else {
		if !slices.Contains(responses.Divisions, input.Division) {
			return nil, huma.Error400BadRequest(fmt.Sprintf("%s isn't a valid division.", input.Division))
		}

		if err := responses.Queries.UpdatePlayerDemoDivision(ctx, queries.UpdatePlayerDemoDivisionParams{
			DemoDivision: sql.NullString{
				String: input.Division,
				Valid:  true,
			},
			ID: input.ID,
		}); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func HandlePostResolvePlayerRequest(ctx context.Context, input *responses.PlayerRequestIDInput) (*struct{}, error) {
	if err := responses.Queries.ResolvePlayerRequest(ctx, input.ID); err != nil {
		return nil, err
	}

	return nil, nil
}

// admin
func HandlePostUpdatePlayerRole(ctx context.Context, input *responses.PlayerRoleInput) (*struct{}, error) {
	if err := responses.Queries.UpdatePlayerRole(ctx, queries.UpdatePlayerRoleParams{
		Role: input.Role,
		ID:   input.ID,
	}); err != nil {
		return nil, err
	}

	return nil, nil
}
