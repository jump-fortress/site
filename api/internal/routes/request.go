package routes

import (
	"context"
	"fmt"
	"slices"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jump-fortress/site/db"
	"github.com/jump-fortress/site/db/queries"
	"github.com/jump-fortress/site/internal/principal"
	"github.com/jump-fortress/site/models"
)

func HandleSubmitRequest(ctx context.Context, input *models.RequestInput) (*struct{}, error) {
	principal, ok := principal.Get(ctx)
	if !ok {
		return nil, models.SessionErr()
	}

	// check for existing pending request
	requestExists, err := db.Queries.CheckPendingRequestExists(ctx, queries.CheckPendingRequestExistsParams{
		PlayerID: principal.SteamID.String(),
		Kind:     input.Kind,
	})
	if err != nil {
		return nil, models.WrapDBErr(err)
	}
	if requestExists == 1 {
		return nil, huma.Error409Conflict(fmt.Sprintf("%s request already exists", input.Kind))
	}

	// validate alias
	if input.Kind == "alias update" {
		if len(input.Content) > 32 {
			return nil, huma.Error400BadRequest("alias is too long (<32 characters)")
		}
		if !AliasRegex.MatchString(input.Content) {
			return nil, huma.Error400BadRequest("alias is invalid (alphanumeric only and in-between spaces, dots, underscores)")
		}
		// div request
	} else {
		player, err := db.Queries.SelectPlayer(ctx, principal.SteamID.String())
		if err != nil {
			return nil, models.WrapDBErr(err)
		}
		if !player.TempusID.Valid {
			return nil, huma.Error400BadRequest("missing a Tempus ID")
		}
		if !slices.Contains(models.Divs, input.Content) || input.Content != "none" {
			return nil, huma.Error400BadRequest(fmt.Sprintf("%s isn't a div", input.Content))
		}
	}

	err = db.Queries.InsertRequest(ctx, queries.InsertRequestParams{
		PlayerID: principal.SteamID.String(),
		Kind:     input.Kind,
		Content:  input.Content,
	})
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	return nil, nil
}

// mod

func HandleResolveRequest(ctx context.Context, input *models.RequestIDInput) (*struct{}, error) {
	err := db.Queries.ResolveRequest(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	return nil, nil
}
