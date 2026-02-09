package routes

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jump-fortress/site/db"
	"github.com/jump-fortress/site/db/queries"
	"github.com/jump-fortress/site/internal/tempus"
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

	// check for duplicates with empty structs
	divs := make(map[string]struct{})
	dupe := false
	for _, ilb := range ilbs {
		_, existingDiv := divs[ilb.Div]
		if existingDiv {
			dupe = true
			break
		}
		divs[ilb.Div] = struct{}{}
	}
	if dupe {
		return nil, huma.Error400BadRequest("can't have duplicate divs")
	}

	// check for divless input with multiple leaderboards
	if len(ilbs) > 1 {
		for _, ilb := range ilbs {
			if ilb.Div == "" {
				return nil, huma.Error400BadRequest("event can't be divless and have multiple divs")
			}
		}
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
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, models.WrapDBErr(err)
		}
	}

	// only events with one leaderboard can be divless
	// can't be divless if event has more than 1 leaderboard, or if it has a single leaderboard with a div
	inputDivless := len(ilbs) == 1 && ilbs[0].Div == ""
	if !inputDivless && !slices.Contains(models.Divs, ilbs[0].Div) {
		return nil, models.DivErr(ilbs[0].Div)
	}
	if len(elbs) > 1 && inputDivless {
		return nil, huma.Error400BadRequest("can't add a divless leaderboard for an event with multiple leaderboards")
	}
	if len(elbs) == 1 && elbs[0].Div.Valid && inputDivless {
		return nil, huma.Error400BadRequest("can't add a divless leaderboard for an event containing a leaderboard with a div")
	}

	// validate input map names
	maps, err := GetMapNames(ctx)
	if err != nil {
		return nil, err
	}

	// if event contains an input div, update it
	// if it doesn't, insert it
	// competition must be deleted to remove leaderboards
	for _, ilb := range ilbs {
		if !slices.Contains(maps, ilb.Map) {
			return nil, huma.Error400BadRequest(fmt.Sprintf("%s isn't a map", ilb.Map))
		}

		contains := false
		for _, elb := range elbs {
			if ilb.ID == elb.ID {
				// update
				err := db.Queries.UpdateLeaderboard(ctx, queries.UpdateLeaderboardParams{
					Div: sql.NullString{
						String: ilb.Div,
						Valid:  !inputDivless,
					},
					Map: ilb.Map,
					ID:  elb.ID,
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

func HandleGetLeaderboardTimes(ctx context.Context, input *models.LeaderboardIDInput) (*models.TimesWithPlayerOutput, error) {
	twps, err := db.Queries.SelectPRTimesFromLeaderboard(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	resp := &models.TimesWithPlayerOutput{
		Body: []models.TimeWithPlayer{},
	}

	for _, twp := range twps {
		resp.Body = append(resp.Body, models.TimeWithPlayer{
			Time:   models.GetTimeResponse(twp.Time),
			Player: models.GetPlayerResponse(twp.Player, true),
			Rank:   twp.TimeRank,
		})
	}
	return resp, nil
}

// dev

func HandleUpdateLeaderboardTempusTimes(ctx context.Context, input *models.LeaderboardIDInput) (*struct{}, error) {
	event, err := db.Queries.SelectEventFromLeaderboardID(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}
	tClassID := 3
	if event.Class == "Demo" {
		tClassID = 4
	}
	leaderboard, err := db.Queries.SelectLeaderboard(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}
	twps, err := db.Queries.SelectPRTimesFromLeaderboard(ctx, leaderboard.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	fmt.Println("found " + strconv.Itoa(len(twps)) + " times")
	for _, twp := range twps {
		if !twp.Time.TempusTimeID.Valid {
			fmt.Printf("attempt: %f by %s\n", twp.Time.Duration, twp.Player.Alias.String)
			time.Sleep(time.Second)
			tt, err := tempus.GetPR(leaderboard.Map, twp.Player.TempusID.Int64, int64(tClassID))
			if err != nil {
				fmt.Println("not found.")
				continue
			}
			tDate := time.UnixMilli(int64(tt.Date) * 1000)
			timeDiff := math.Abs(twp.Time.Duration - tt.Duration)
			// assume margins of error
			if timeDiff < 0.015 && tDate.Add(time.Hour*24).After(event.StartsAt) && tDate.Before(event.EndsAt.Add(time.Hour*24)) {
				fmt.Println("found valid.")
				err = db.Queries.UpdateTimeFromTempus(ctx, queries.UpdateTimeFromTempusParams{
					Duration: tt.Duration,
					TempusTimeID: sql.NullInt64{
						Int64: tt.ID,
						Valid: true,
					},
					CreatedAt: tDate,
					ID:        twp.Time.ID,
				})
				if err != nil {
					return nil, models.WrapDBErr(err)
				}
			} else {
				fmt.Println("not during event.")
			}
		}
	}

	return nil, nil
}
