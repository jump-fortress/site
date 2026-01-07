package internal

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"

	"github.com/danielgtaylor/huma/v2"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

func getMapResponse(dbMap queries.Map) responses.Map {
	return responses.Map{
		ID:            dbMap.ID,
		Name:          dbMap.Name,
		Courses:       dbMap.Courses.Int64,
		Bonuses:       dbMap.Bonuses.Int64,
		SoldierTier:   dbMap.SoldierTier,
		DemoTier:      dbMap.DemoTier,
		SoldierRating: dbMap.SoldierRating,
		DemoRating:    dbMap.DemoRating,
	}
}

func getTempusMaps() (*[]responses.TempusMapInfo, error) {
	url := "https://tempus2.xyz/api/v0/maps/detailedList"

	response, err := retryablehttp.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	tempusResponseMaps := &[]responses.TempusMapInfo{}
	if err := json.Unmarshal(body, &tempusResponseMaps); err != nil {
		return nil, err
	}

	return tempusResponseMaps, nil
}

func HandlePostUpdateMaps(ctx context.Context, _ *struct{}) (*struct{}, error) {
	tempusMaps, err := getTempusMaps()
	if err != nil {
		return nil, huma.Error503ServiceUnavailable("couldn't verify with Tempus. if Tempus isn't down, please check your Tempus ID again")
	}

	for _, m := range *tempusMaps {
		hasCourses := m.Zones.Bonus != 0
		hasBonuses := m.Zones.Course != 0
		if err := responses.Queries.InsertMap(ctx, queries.InsertMapParams{
			ID:   m.ID,
			Name: m.Name,
			Courses: sql.NullInt64{
				Int64: m.Zones.Course,
				Valid: hasCourses,
			},
			Bonuses: sql.NullInt64{
				Int64: m.Zones.Bonus,
				Valid: hasBonuses,
			},
			SoldierTier:   m.Tier.Soldier,
			DemoTier:      m.Tier.Demo,
			SoldierRating: m.Rating.Soldier,
			DemoRating:    m.Rating.Demo,
		}); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func HandleGetAllMaps(ctx context.Context, _ *struct{}) (*responses.MapsOutput, error) {
	maps, err := responses.Queries.GetMaps(ctx)
	if err != nil {
		return nil, err
	}

	resp := &responses.MapsOutput{
		Body: []responses.Map{},
	}

	for _, m := range maps {
		mapResponse := getMapResponse(m)
		resp.Body = append(resp.Body, mapResponse)
	}

	return resp, nil
}

func HandleGetAllMapNames(ctx context.Context, _ *struct{}) (*responses.MapNamesOutput, error) {
	maps, err := responses.Queries.GetMapNames(ctx)
	if err != nil {
		return nil, err
	}

	resp := &responses.MapNamesOutput{
		Body: maps,
	}

	return resp, nil
}
