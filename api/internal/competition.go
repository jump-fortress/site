package internal

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

func getCompetitionID(ctx context.Context, competitionType string, typeID int64) (int64, error) {
	var err error = nil
	var id int64 = 0

	switch competitionType {
	case "Monthly":
		id, err = responses.Queries.SelectMonthlyCompetitionID(ctx, typeID)
		// case "Motw":
		// 	return id, huma.Error404NotFound("todo not implemented")
		//
		// case "Bounty":
		// 	return id, huma.Error404NotFound("todo not implemented")
		//
		// case "Quest":
		// 	return id, huma.Error404NotFound("todo not implemented")
		//
		// case "Playoffs":
		// 	return id, huma.Error404NotFound("todo not implemented")
		//
		// case "JWC":
		// 	return id, huma.Error404NotFound("todo not implemented")
		//
		// case "Archive":
		// 	return id, huma.Error404NotFound("todo not implemented")
		//
		// case "Other":
		// 	return id, huma.Error404NotFound("todo not implemented")
	}

	return id, err
}

func HandlePostCancelCompetition(ctx context.Context, input *responses.CompetitionIDInput) (*struct{}, error) {
	// this query also checks that the competition hasn't started yet
	competition, err := responses.Queries.DeleteCompetition(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	jsonCompetition, err := json.Marshal(competition)
	if err != nil {
		return nil, err
	}

	if err := responses.Queries.InsertDeletedRecord(ctx, queries.InsertDeletedRecordParams{
		SourceTable: "competition",
		SourceID:    strconv.FormatInt(competition.ID, 10),
		Data:        json.RawMessage(jsonCompetition),
	}); err != nil {
		return nil, err
	}

	return nil, nil
}
