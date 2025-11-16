package internal

import (
	"context"
	"strconv"

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
