package internal

import (
	"context"
	"strconv"

	"github.com/spiritov/jump/api/db/responses"
)

func HandleGetPlayer(ctx context.Context, input *responses.PlayerIDInput) (*responses.PlayerOutput, error) {
	player, err := responses.Queries.SelectPlayer(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayerOutput{Body: player}
	return resp, nil
}

func HandleGetPlayerBySteamID64(ctx context.Context, input *responses.PlayerSteamID64Input) (*responses.PlayerOutput, error) {
	steamID64_string := strconv.FormatUint(input.SteamID64, 10)

	player, err := responses.Queries.SelectPlayerFromSteamID64(ctx, steamID64_string)
	if err != nil {
		return nil, err
	}

	resp := &responses.PlayerOutput{Body: player}
	return resp, nil
}
