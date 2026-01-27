package routes

import (
	"context"
	"database/sql"

	"github.com/jump-fortress/site/db"
	"github.com/jump-fortress/site/db/queries"
	"github.com/jump-fortress/site/internal/tempus"
	"github.com/jump-fortress/site/models"
)

func HandleUpdateMaps(ctx context.Context, input *struct{}) (*struct{}, error) {
	// get Tempus map info
	tms, err := tempus.GetMaps()
	if err != nil {
		return nil, models.WrapTempusErr(err)
	}

	for _, tm := range *tms {
		// check for bonus / courses first
		hasBonus := tm.Zones.Bonus != 0
		hasCourse := tm.Zones.Course != 0

		err := db.Queries.InsertMap(ctx, queries.InsertMapParams{
			Name: tm.Name,
			Courses: sql.NullInt64{
				Int64: tm.Zones.Course,
				Valid: hasCourse,
			},
			Bonuses: sql.NullInt64{
				Int64: tm.Zones.Bonus,
				Valid: hasBonus,
			},
			SoldierTier:   tm.Tier.Soldier,
			DemoTier:      tm.Tier.Demo,
			SoldierRating: tm.Rating.Soldier,
			DemoRating:    tm.Rating.Demo,
		})
		if err != nil {
			return nil, models.WrapDBErr(err)
		}
	}

	return nil, nil
}

func HandleGetMaps(ctx context.Context, input *struct{}) (*models.MapsOutput, error) {
	maps, err := db.Queries.SelectMaps(ctx)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	resp := &models.MapsOutput{
		Body: []models.Map{},
	}
	for _, m := range maps {
		resp.Body = append(resp.Body, models.GetMapResponse(m))
	}
	return resp, nil
}
