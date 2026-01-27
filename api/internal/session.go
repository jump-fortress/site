package internal

import (
	"context"
	"database/sql"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jump-fortress/site/db"
	"github.com/jump-fortress/site/db/queries"
	"github.com/jump-fortress/site/internal/principal"
	"github.com/jump-fortress/site/internal/routes"
	"github.com/jump-fortress/site/models"
)

func HandleGetSession(ctx context.Context, _ *struct{}) (*models.SessionOutput, error) {
	principal, ok := principal.Get(ctx)
	if !ok {
		return nil, models.SessionErr()
	}

	player, err := db.Queries.SelectPlayer(ctx, principal.SteamID.String())
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	// set alias or avatar from Steam info, if either are missing
	if !player.Alias.Valid || !player.AvatarUrl.Valid {
		steamProfile, err := FetchProfileSummary(principal.SteamID.ID())
		if err != nil {
			return nil, huma.Error503ServiceUnavailable("something went wrong getting this Steam profile")
		}

		if !player.Alias.Valid {
			steamAlias := steamProfile.PersonaName

			// set default alias if Steam name doesn't match
			if !routes.AliasRegex.MatchString(steamAlias) {
				steamAlias = "update alias"
			}

			// then update it
			err := db.Queries.UpdatePlayerAlias(ctx, queries.UpdatePlayerAliasParams{
				Alias: sql.NullString{
					String: steamAlias,
					Valid:  true,
				},
				ID: player.ID,
			})
			if err != nil {
				return nil, models.WrapDBErr(err)
			}
			player.Alias = sql.NullString{
				String: steamAlias,
				Valid:  true,
			}
		}

		if !player.AvatarUrl.Valid {
			err := db.Queries.UpdatePlayerAvatarURL(ctx, queries.UpdatePlayerAvatarURLParams{
				AvatarUrl: sql.NullString{
					String: steamProfile.AvatarFullURL,
					Valid:  true,
				},
				ID: player.ID,
			})
			if err != nil {
				return nil, models.WrapDBErr(err)
			}
			player.AvatarUrl = sql.NullString{
				String: steamProfile.AvatarFullURL,
				Valid:  true,
			}
		}
	}

	resp := &models.SessionOutput{Body: models.GetSessionResponse(player)}
	return resp, nil
}
