package internal

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

func updateCompetitionPrizepool(ctx context.Context, id int64) error {
	total, err := responses.Queries.SumCompetitionPrizepool(ctx, id)
	if err != nil {
		return err
	}

	if err := responses.Queries.UpdateCompetitionPrizepool(ctx, queries.UpdateCompetitionPrizepoolParams{
		Prizepool: sql.NullInt64{
			Int64: total,
			Valid: true,
		},
		ID: id,
	}); err != nil {
		return err
	}

	return nil
}

func getPrizesResponse(cdPrizepool []queries.SelectDivisionPrizepoolRow) []responses.CompetitionPrize {
	rcp := []responses.CompetitionPrize{}

	for _, cdp := range cdPrizepool {
		rcp = append(rcp, responses.CompetitionPrize{
			ID:         cdp.ID,
			DivisionID: cdp.CompetitionDivisionID,
			Placement:  cdp.Placement,
			Amount:     cdp.Amount,
		})
	}

	return rcp
}

func getCdPrizepoolResponse(cdPrizepool []queries.SelectDivisionPrizepoolRow, cd *responses.CompetitionDivision) responses.DivisionPrizepool {
	prizes := getPrizesResponse(cdPrizepool)

	return responses.DivisionPrizepool{
		Division: responses.CompetitionDivision{
			ID:            cd.ID,
			CompetitionID: cd.CompetitionID,
			Division:      cd.Division,
			Map:           cd.Map,
		},
		Prizes: prizes,
	}
}

// open

func HandlePostGetPrizepool(ctx context.Context, input *responses.CompetitionIDInput) (*responses.PrizepoolOutput, error) {
	cds, err := responses.Queries.SelectCompetitionDivisions(ctx, input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("competition not found with id " + strconv.FormatInt(input.ID, 10))
	}

	prizepool := &responses.PrizepoolOutput{
		Body: []responses.DivisionPrizepool{},
	}

	for _, cd := range cds {
		cdPrizepool, err := responses.Queries.SelectDivisionPrizepool(ctx, cd.ID)
		if err != nil {
			return nil, huma.Error500InternalServerError("error getting division prizepools")
		}

		cdPrizepoolResponse := getCdPrizepoolResponse(cdPrizepool, &responses.CompetitionDivision{
			ID:            cd.ID,
			CompetitionID: cd.CompetitionID,
			Division:      cd.Division,
			Map:           cd.Map,
		})
		prizepool.Body = append(prizepool.Body, cdPrizepoolResponse)
	}

	return prizepool, nil
}

// admin

func HandlePostCreatePrizepool(ctx context.Context, input *responses.DivisionPrizepoolInput) (*struct{}, error) {
	cd, err := responses.Queries.SelectCompetitionDivision(ctx, input.Body.DivisionID)
	if err != nil {
		return nil, huma.Error400BadRequest("no competition division found with id " + strconv.FormatInt(input.Body.DivisionID, 10))
	}

	competition, err := responses.Queries.SelectCompetition(ctx, cd.CompetitionID)
	if err != nil {
		return nil, huma.Error400BadRequest("no competition found with id " + strconv.FormatInt(cd.CompetitionID, 10))
	}

	now := time.Now()
	if competition.StartsAt.Before(now) {
		return nil, huma.Error400BadRequest("competition already started")
	}

	// make sure division prizepool is empty
	dp, err := responses.Queries.SelectDivisionPrizepool(ctx, input.Body.DivisionID)
	if len(dp) != 0 {
		return nil, huma.Error400BadRequest("division already has a prizepool")
	}

	// insert prize per input to prizepool
	for _, prize := range input.Body.Prizepool.Prizes {
		if err := responses.Queries.InsertCompetitionPrize(ctx, queries.InsertCompetitionPrizeParams{
			CompetitionDivisionID: prize.DivisionID,
			Placement:             prize.Placement,
			Amount:                prize.Amount,
		}); err != nil {
			return nil, huma.Error500InternalServerError("error creating prizepool")
		}
	}

	if err := updateCompetitionPrizepool(ctx, competition.ID); err != nil {
		return nil, err
	}

	return nil, nil
}

func HandlePostDeletePrizepool(ctx context.Context, input *responses.DivisionIDInput) (*struct{}, error) {
	cd, err := responses.Queries.SelectCompetitionDivision(ctx, input.DivisionID)
	if err != nil {
		return nil, huma.Error400BadRequest("no competition division found with id " + strconv.FormatInt(input.DivisionID, 10))
	}

	competition, err := responses.Queries.SelectCompetition(ctx, cd.CompetitionID)
	if err != nil {
		return nil, huma.Error400BadRequest("no competition found with id " + strconv.FormatInt(cd.CompetitionID, 10))
	}

	now := time.Now()
	if competition.StartsAt.Before(now) {
		return nil, huma.Error400BadRequest("competition already started")
	}

	dp, err := responses.Queries.SelectDivisionPrizepool(ctx, input.DivisionID)
	if len(dp) == 0 {
		return nil, huma.Error400BadRequest("division has no prizepool")
	}

	if err := responses.Queries.DeleteCompetitionDivisionPrizepool(ctx, cd.ID); err != nil {
		return nil, huma.Error500InternalServerError("error")
	}

	if err := updateCompetitionPrizepool(ctx, competition.ID); err != nil {
		return nil, err
	}

	return nil, nil
}
