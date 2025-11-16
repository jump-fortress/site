package internal

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

// todo: error handling / logging

// map a database player to the session fields we want
func SessionFromPlayer(player queries.Player) responses.Session {
	return responses.Session{
		ID:             player.ID,
		DisplayName:    player.DisplayName.String,
		SteamAvatarURL: player.SteamAvatarUrl.String,
	}
}

func HandleGetSession(ctx context.Context, _ *struct{}) (*responses.SessionOutput, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	// use steamID64 to derive a session from a player in the database
	steamID64_string := strconv.FormatUint(principal.SteamID, 10)
	player, err := responses.Queries.SelectPlayerFromSteamID64(ctx, steamID64_string)
	if err != nil {
		return nil, err
	}

	var session responses.Session

	// if this player is missing fields, set these required fields with their Steam info
	if !player.DisplayName.Valid || !player.SteamAvatarUrl.Valid {
		steamProfileSummary, err := FetchProfileSummary(principal.SteamID)
		if err != nil {
			return nil, err
		}

		updatedPlayer, err := responses.Queries.UpdatePlayerSessionInfo(ctx, queries.UpdatePlayerSessionInfoParams{
			SteamAvatarUrl: sql.NullString{
				String: steamProfileSummary.AvatarFullURL,
				Valid:  true,
			},
			DisplayName: sql.NullString{
				String: steamProfileSummary.PersonaName,
				Valid:  true,
			},
			SteamId64: steamID64_string,
		})
		if err != nil {
			return nil, err
		}

		session = SessionFromPlayer(updatedPlayer)
	} else {
		session = SessionFromPlayer(player)
	}
	resp := &responses.SessionOutput{Body: session}
	return resp, nil
}
