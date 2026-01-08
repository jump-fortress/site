package internal

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

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
