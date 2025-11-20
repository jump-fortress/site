package internal

import (
	"context"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

// todo: check some nullable before using them?
func PlayerResponseFromPlayer(player queries.Player) responses.Player {
	return responses.Player{
		ID:              player.ID,
		Role:            player.Role,
		SteamAvatarUrl:  player.SteamAvatarUrl.String,
		TempusID:        player.TempusID.Int64,
		DisplayName:     player.DisplayName.String,
		SoldierDivision: player.SoldierDivision.String,
		DemoDivision:    player.DemoDivision.String,
		PreferredClass:  player.PreferredClass,
		CreatedAt:       player.CreatedAt,
	}
}

func FullPlayerResponseFromPlayer(player queries.Player) responses.FullPlayer {
	return responses.FullPlayer{
		ID:              player.ID,
		Role:            player.Role,
		SteamId64:       player.SteamId64,
		SteamAvatarUrl:  player.SteamAvatarUrl.String,
		SteamTradeToken: player.SteamTradeToken.String,
		TempusID:        player.TempusID.Int64,
		DiscordID:       player.DiscordID.String,
		DisplayName:     player.DisplayName.String,
		SoldierDivision: player.SoldierDivision.String,
		DemoDivision:    player.DemoDivision.String,
		PreferredClass:  player.PreferredClass,
		CreatedAt:       player.CreatedAt,
	}
}

func HandleGetPlayer(ctx context.Context, input *responses.PlayerIDInput) (*responses.PlayerOutput, error) {
	player, err := responses.Queries.SelectPlayer(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayerOutput{Body: PlayerResponseFromPlayer(player)}
	return resp, nil
}

func HandleGetPlayerBySteamID64(ctx context.Context, input *responses.PlayerSteamID64Input) (*responses.PlayerOutput, error) {
	steamID64_string := strconv.FormatUint(input.SteamID64, 10)

	player, err := responses.Queries.SelectPlayerFromSteamID64(ctx, steamID64_string)
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayerOutput{Body: PlayerResponseFromPlayer(player)}
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
			Player: PlayerResponseFromPlayer(player),
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

func HandleGetAllPlayers(ctx context.Context, _ *struct{}) (*responses.ManyPlayersOutput, error) {
	players, err := responses.Queries.SelectAllPlayers(ctx)
	if err != nil {
		return nil, err
	}

	resp := &responses.ManyPlayersOutput{
		Body: []responses.Player{},
	}

	for _, p := range players {
		playerResponse := PlayerResponseFromPlayer(p)
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
		fullPlayerResponse := FullPlayerResponseFromPlayer(p)
		resp.Body = append(resp.Body, fullPlayerResponse)
	}

	return resp, nil
}
