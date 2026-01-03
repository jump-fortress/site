package internal

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

func getMonthlyResponse(monthly queries.SelectAllMonthlyRow) responses.Monthly {
	return responses.Monthly{
		Competition: responses.Competition{
			ID:        monthly.ID,
			Class:     monthly.Class,
			StartsAt:  monthly.StartsAt,
			EndsAt:    monthly.EndsAt,
			CreatedAt: monthly.CreatedAt,
		},
		Divisions: []responses.CompetitionDivision{},
	}
}

func HandlePostMonthly(ctx context.Context, input *responses.MonthlyInput) (*struct{}, error) {
	competition, err := responses.Queries.InsertCompetition(ctx, queries.InsertCompetitionParams{
		Class:    input.Body.Competition.Class,
		StartsAt: input.Body.Competition.StartsAt,
		EndsAt:   input.Body.Competition.StartsAt.AddDate(0, 0, 2),
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

func HandleGetMonthly(ctx context.Context, _ *struct{}) (*responses.MonthliesOutput, error) {
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

	if len(resp.Body) == 0 {
		return nil, huma.Error404NotFound("no monthlies found")
	}
	return resp, nil
}
