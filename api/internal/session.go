package internal

import (
	"context"
	"database/sql"

	"github.com/danielgtaylor/huma/v2"
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

// todo: error handling / logging

// map a database player to the session fields we want
func SessionFromPlayer(player queries.Player) responses.Session {
	return responses.Session{
		ID:             player.ID,
		Role:           player.Role,
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
	player, err := responses.Queries.SelectPlayer(ctx, principal.SteamID.String())
	if err != nil {
		return nil, err
	}

	var session responses.Session

	// only fetch from steam if something is needed
	if !player.DisplayName.Valid || !player.SteamAvatarUrl.Valid {
		steamProfileSummary, err := FetchProfileSummary(principal.SteamID.ID())
		if err != nil {
			return nil, err
		}

		// set initial display name
		if !player.DisplayName.Valid {
			var displayName = steamProfileSummary.PersonaName
			if !displayNameRegex.MatchString(displayName) {
				displayName = "invalid alias" // or something..
			}

			if err := responses.Queries.UpdatePlayerDisplayName(ctx, queries.UpdatePlayerDisplayNameParams{
				DisplayName: sql.NullString{
					String: displayName,
					Valid:  true,
				},
				ID: player.ID,
			}); err != nil {
				return nil, err
			}

			player.DisplayName = sql.NullString{
				String: displayName,
				Valid:  true,
			}
		}

		// set initial avatar
		if !player.SteamAvatarUrl.Valid {
			if err := responses.Queries.UpdatePlayerSteamAvatarURL(ctx, queries.UpdatePlayerSteamAvatarURLParams{
				SteamAvatarUrl: sql.NullString{
					String: steamProfileSummary.AvatarFullURL,
					Valid:  true,
				},
				ID: player.ID,
			}); err != nil {
				return nil, err
			}

			player.SteamAvatarUrl = sql.NullString{
				String: steamProfileSummary.AvatarFullURL,
				Valid:  true,
			}
		}
	}

	session = SessionFromPlayer(player)
	resp := &responses.SessionOutput{Body: session}

	return resp, nil
}
