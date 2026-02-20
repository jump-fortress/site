package routes

import (
	"context"

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

	// todo: validate kind?
	// content can be empty when requesting no division
	err := db.Queries.InsertRequest(ctx, queries.InsertRequestParams{
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
