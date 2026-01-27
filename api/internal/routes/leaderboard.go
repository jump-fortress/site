package routes

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jump-fortress/site/db"
	"github.com/jump-fortress/site/db/queries"
	"github.com/jump-fortress/site/models"
)

func GetEventAndLeaderboard(ctx context.Context, leaderboard_id int64) (*queries.Event, *queries.Leaderboard, error) {
	leaderboard, err := db.Queries.SelectLeaderboard(ctx, leaderboard_id)
	if err != nil {
		return nil, nil, models.WrapDBErr(err)
	}
	event, err := db.Queries.SelectEventFromLeaderboardID(ctx, leaderboard.ID)
	if err != nil {
		return nil, nil, models.WrapDBErr(err)
	}

	return &event, &leaderboard, nil
}

func HandleUpdateLeaderboards(ctx context.Context, input *models.LeaderboardsInput) (*struct{}, error) {
	// validate leaderboard info
	ilbs := input.Body
	if len(ilbs) == 0 {
		return nil, huma.Error400BadRequest("event must have at least one leaderboard")
	}

	eventID := ilbs[0].EventID
	event, err := db.Queries.SelectEvent(ctx, eventID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	// if event has started, leaderboards can't be changed
	now := time.Now()
	if event.StartsAt.Before(now) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("event has already started (%s)", event.StartsAt))
	}

	elbs, err := db.Queries.SelectLeaderboards(ctx, eventID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}
	inputDivless := len(ilbs) == 1 && ilbs[0].Div == ""

	// only events with one leaderboard can be divless
	// can't be divless if event has more than 1 leaderboard, or if it has a single leaderboard with a div
	if !inputDivless || !slices.Contains(models.Divs, ilbs[0].Div) {
		return nil, models.DivErr(ilbs[0].Div)
	}
	if len(elbs) > 1 && inputDivless {
		return nil, huma.Error400BadRequest("can't add a divisionless leaderboard for an event with multiple leaderboards")
	}
	if len(elbs) == 1 && elbs[0].Div.Valid && inputDivless {
		return nil, huma.Error400BadRequest("can't add a divisionless leaderboard for an event containing a leaderboard with a division")
	}

	// if event contains an input div, update it
	// if it doesn't, insert it
	// competition must be deleted to remove leaderboards
	for _, ilb := range ilbs {
		contains := false
		for _, elb := range elbs {
			if ilb.Div == elb.Div.String {
				// update
				err := db.Queries.UpdateLeaderboard(ctx, queries.UpdateLeaderboardParams{
					Div: sql.NullString{
						String: ilb.Div,
						Valid:  !inputDivless,
					},
					Map: ilb.Map,
					ID:  ilb.ID,
				})
				if err != nil {
					return nil, models.WrapDBErr(err)
				}
				contains = true
				break
			}
		}
		if !contains {
			// add
			err := db.Queries.InsertLeaderboard(ctx, queries.InsertLeaderboardParams{
				EventID: eventID,
				Div: sql.NullString{
					String: ilb.Div,
					Valid:  !inputDivless,
				},
				Map: ilb.Map,
			})
			if err != nil {
				return nil, models.WrapDBErr(err)
			}
		}
	}

	return nil, nil
}

func HandleGetLeaderboardTimes(ctx context.Context, input *models.LeaderboardIDInput) (*models.TimesOutput, error) {
	times, err := db.Queries.SelectTimesFromLeaderboard(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	resp := &models.TimesOutput{
		Body: []models.Time{},
	}

	for _, t := range times {
		resp.Body = append(resp.Body, models.GetTimeResponse(t))
	}
	return resp, nil
}
