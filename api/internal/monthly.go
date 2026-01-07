package internal

import (
	"context"
	"slices"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

// todo: monthly id

func getMonthlyResponse(monthly queries.SelectAllMonthlyRow) responses.Monthly {
	return responses.Monthly{
		ID: monthly.CompetitionID,
		Competition: responses.Competition{
			ID:        monthly.ID,
			Class:     monthly.Class,
			StartsAt:  monthly.StartsAt,
			EndsAt:    monthly.EndsAt,
			VisibleAt: monthly.VisibleAt,
			Complete:  monthly.Complete,
			CreatedAt: monthly.CreatedAt,
		},
		Divisions: []responses.CompetitionDivision{},
	}
}

func HandleGetAllMonthlies(ctx context.Context, _ *struct{}) (*responses.MonthliesOutput, error) {
	monthlies, err := responses.Queries.SelectAllMonthly(ctx)
	if err != nil {
		return nil, err
	}

	resp := &responses.MonthliesOutput{
		Body: []responses.Monthly{},
	}

	for _, m := range monthlies {
		now := time.Now()

		// skip this monthly if it's not visible yet
		if m.VisibleAt.After(now) {
			continue
		}

		monthlyResponse := getMonthlyResponse(m)

		divisions, err := responses.Queries.SelectCompetitionDivisions(ctx, m.CompetitionID)
		if err != nil {
			return nil, err
		}

		for _, d := range divisions {
			// hide maps if competition hasn't started
			if m.StartsAt.Before(now) {
				d.Map = ""
			}
			monthlyResponse.Divisions = append(monthlyResponse.Divisions, getCompetitionDivisionResponse(d))
		}

		resp.Body = append(resp.Body, monthlyResponse)
	}

	return resp, nil
}

// admin

func HandlePostCreateMonthly(ctx context.Context, input *responses.MonthlyInput) (*struct{}, error) {
	maps, err := responses.Queries.GetMapNames(ctx)
	if err != nil {
		return nil, err
	}

	for _, cd := range input.Body.Divisions {
		if !slices.Contains(maps, cd.Map) {
			return nil, huma.Error400BadRequest("invalid map name")
		}
	}

	if input.Body.Competition.StartsAt.Before(input.Body.Competition.VisibleAt) {
		return nil, huma.Error400BadRequest("competition must be visible before it starts")
	}

	now := time.Now()
	if input.Body.Competition.StartsAt.Before(now) || input.Body.Competition.VisibleAt.Before(now) {
		return nil, huma.Error400BadRequest("competition must start and be visible in the future")
	}

	competition, err := responses.Queries.InsertCompetition(ctx, queries.InsertCompetitionParams{
		Class:     input.Body.Competition.Class,
		StartsAt:  input.Body.Competition.StartsAt,
		EndsAt:    input.Body.Competition.StartsAt.AddDate(0, 0, 2), // 2 days
		VisibleAt: input.Body.Competition.VisibleAt,
	})
	if err != nil {
		return nil, err
	}

	// create monthly
	if err := responses.Queries.InsertMonthly(ctx, competition.ID); err != nil {
		return nil, err
	}

	// create competition
	for _, d := range input.Body.Divisions {
		if err := responses.Queries.InsertCompetitionDivision(ctx, queries.InsertCompetitionDivisionParams{
			CompetitionID: competition.ID,
			Division:      d.Division,
			Map:           d.Map,
		}); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func HandleGetAllFullMonthlies(ctx context.Context, _ *struct{}) (*responses.MonthliesOutput, error) {
	monthlies, err := responses.Queries.SelectAllMonthly(ctx)
	if err != nil {
		return nil, err
	}

	resp := &responses.MonthliesOutput{
		Body: []responses.Monthly{},
	}

	for _, m := range monthlies {
		monthlyResponse := getMonthlyResponse(m)

		divisions, err := responses.Queries.SelectCompetitionDivisions(ctx, m.CompetitionID)
		if err != nil {
			return nil, err
		}

		for _, d := range divisions {
			monthlyResponse.Divisions = append(monthlyResponse.Divisions, getCompetitionDivisionResponse(d))
		}

		resp.Body = append(resp.Body, monthlyResponse)
	}

	return resp, nil
}
