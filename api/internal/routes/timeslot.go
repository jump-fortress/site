package routes

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jump-fortress/site/db"
	"github.com/jump-fortress/site/db/queries"
	"github.com/jump-fortress/site/internal/principal"
	"github.com/jump-fortress/site/models"
)

var (
	MotwDuration = time.Hour * 3
)

func GetTimeslotDatetimes(timeslot queries.MotwTimeslot, date time.Time) models.TimeslotDatetimes {
	starts := time.Date(date.Year(), date.Month(), date.Day(), timeslot.StartsAt.Hour(), timeslot.StartsAt.Minute(), timeslot.StartsAt.Second(), timeslot.StartsAt.Nanosecond(), time.UTC)
	ends := starts.Add(MotwDuration)
	return models.TimeslotDatetimes{
		ID:       timeslot.ID,
		StartsAt: starts,
		EndsAt:   ends,
	}
}

func HandleUpdateTimeslotPref(ctx context.Context, input *models.TimeslotIDInput) (*struct{}, error) {
	principal, ok := principal.Get(ctx)
	if !ok {
		return nil, models.SessionErr()
	}

	// validate that an motw isn't upcoming first
	recentMotw, err := db.Queries.SelectLastMOTW(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, models.WrapDBErr(err)
	}

	now := time.Now().UTC()
	if recentMotw.StartsAt.Before(now) && recentMotw.EndsAt.After(now) {
		return nil, huma.Error400BadRequest("can't update timeslot while a motw is in progress")
	}

	err = db.Queries.UpdatePlayerTimeslot(ctx, queries.UpdatePlayerTimeslotParams{
		TimeslotID: input.ID,
		PlayerID:   principal.SteamID.String(),
	})
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	return nil, nil
}

func HandleGetTimeslotInfo(ctx context.Context, input *models.EventIDInput) (*models.TimeslotInfoOutput, error) {
	principal, ok := principal.Get(ctx)
	if !ok {
		return nil, models.SessionErr()
	}

	timeslots, err := db.Queries.SelectTimeslots(ctx)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}
	playerTimeslot, err := db.Queries.SelectPlayerTimeslot(ctx, principal.SteamID.String())
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	resp := &models.TimeslotInfoOutput{
		Body: models.TimeslotInfo{
			Timeslots: []models.TimeslotDatetimes{},
			PlayerTimeslot: models.PlayerTimeslot{
				TimeslotID: playerTimeslot.MotwTimeslot.ID,
				PlayerID:   playerTimeslot.PlayerMotwTimeslot.PlayerID,
			},
		},
	}

	// use event start date for timeslots, unless event id was 0
	timeslotDate := time.Now().UTC()
	if input.ID != 0 {
		event, err := db.Queries.SelectEvent(ctx, input.ID)
		if err != nil {
			return nil, huma.Error404NotFound("event not found")
		}
		timeslotDate = event.StartsAt
	}

	for _, t := range timeslots {
		ets := GetTimeslotDatetimes(t, timeslotDate)
		resp.Body.Timeslots = append(resp.Body.Timeslots, models.TimeslotDatetimes{
			ID:       ets.ID,
			StartsAt: ets.StartsAt,
			EndsAt:   ets.EndsAt,
		})
	}
	return resp, nil
}

// admin

func HandleUpdateTimeslot(ctx context.Context, input *models.TimeslotInput) (*struct{}, error) {
	// validate that an motw isn't upcoming first
	recentMotw, err := db.Queries.SelectLastMOTW(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, models.WrapDBErr(err)
	}

	now := time.Now().UTC()
	if err == nil {
		if recentMotw.EndsAt.After(now) {
			return nil, huma.Error400BadRequest("can't update timeslots while a motw is upcoming")
		}
	}

	its := input.Body

	inputTimeslot := GetTimeslotDatetimes(queries.MotwTimeslot{
		ID:       its.ID,
		StartsAt: its.StartsAt,
	}, now)

	// validate input timeslot doesn't overlap with other timeslots, if any exist
	timeslots, err := db.Queries.SelectTimeslots(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, models.WrapDBErr(err)
	}

	for _, ts := range timeslots {
		// skip checking same ID
		if ts.ID == input.Body.ID {
			continue
		}
		cts := GetTimeslotDatetimes(ts, now)
		if its.ID < ts.ID && inputTimeslot.StartsAt.After(cts.StartsAt) {
			return nil, huma.Error400BadRequest("input timeslot can't start after a timeslot with a higher ID.")
		}
		if its.ID > ts.ID && inputTimeslot.StartsAt.Before(cts.StartsAt) {
			return nil, huma.Error400BadRequest("input timeslot can't start before a timeslot with a lower ID.")
		}
		// input timeslot starts or ends during another
		if inputTimeslot.StartsAt.After(cts.StartsAt) && inputTimeslot.StartsAt.Before(cts.EndsAt) || inputTimeslot.EndsAt.After(cts.StartsAt) && inputTimeslot.EndsAt.Before(cts.EndsAt) {
			return nil, huma.Error400BadRequest("input timeslot overlaps with an existing timeslot")
		}
	}

	err = db.Queries.UpsertTimeslot(ctx, queries.UpsertTimeslotParams{
		ID:       input.Body.ID,
		StartsAt: inputTimeslot.StartsAt,
	})
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	return nil, nil
}
