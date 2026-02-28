package routes

import (
	"context"

	"github.com/jump-fortress/site/db"
	"github.com/jump-fortress/site/db/queries"
	"github.com/jump-fortress/site/models"
)

func HandleGetPrizepoolTotal(ctx context.Context, input *models.EventIDInput) (*models.PrizepoolTotalOutput, error) {
	total, err := db.Queries.SelectPrizepoolTotal(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}
	return &models.PrizepoolTotalOutput{
		Body: models.PrizepoolTotal{
			Total: total,
		},
	}, nil
}

func HandleGetLeaderboardPrizepool(ctx context.Context, input *models.LeaderboardIDInput) (*models.PrizepoolOutput, error) {
	prizepool, err := db.Queries.SelectPrizepool(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	resp := &models.PrizepoolOutput{
		Body: []models.Prize{},
	}
	for _, p := range prizepool {
		resp.Body = append(resp.Body, models.GetPrizeResponse(p))
	}
	return resp, nil
}

// admin

func HandleUpdateLeaderboardPrizepool(ctx context.Context, input *models.PrizepoolInput) (*struct{}, error) {
	// validate event hasn't ended
	_, err := db.Queries.SelectEventFromLeaderboardID(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}
	// todo: remove bypass
	// now := time.Now().UTC()
	// if event.EndsAt.Before(now) {
	// 	return nil, huma.Error400BadRequest(fmt.Sprintf("event has ended (%s)", event.EndsAt))
	// }

	// delete existing prizepool
	err = db.Queries.DeletePrizepool(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	// insert new prizepool
	for _, p := range input.Body {
		err := db.Queries.InsertPrize(ctx, queries.InsertPrizeParams{
			LeaderboardID: p.LeaderboardID,
			Position:      p.Position,
			Keys:          p.Keys,
		})
		if err != nil {
			return nil, models.WrapDBErr(err)
		}
	}

	return nil, nil
}
