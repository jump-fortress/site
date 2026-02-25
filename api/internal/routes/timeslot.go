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

func GetTimeslotDatetimes(timeslot queries.MotwTimeslot, date time.Time) (time.Time, time.Time) {
	starts := time.Date(date.Year(), date.Month(), date.Day(), timeslot.StartsAt.Hour(), timeslot.StartsAt.Minute(), timeslot.StartsAt.Second(), timeslot.StartsAt.Nanosecond(), time.UTC)
	ends := starts.Add(MotwDuration)
	return starts, ends
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

func HandleGetTimeslot(ctx context.Context, _ *struct{}) (*models.TimeslotInfoOutput, error) {
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
			Timeslots: []models.MOTWTimeslot{},
			PlayerTimeslot: models.PlayerTimeslot{
				TimeslotID: playerTimeslot.MotwTimeslot.ID,
				PlayerID:   playerTimeslot.PlayerMotwTimeslot.PlayerID,
			},
		},
	}
	for _, t := range timeslots {
		resp.Body.Timeslots = append(resp.Body.Timeslots, models.MOTWTimeslot{
			ID:       t.ID,
			StartsAt: t.StartsAt,
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
	itsStarts, itsEnds := GetTimeslotDatetimes(queries.MotwTimeslot(its), now)

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
		tsStarts, tsEnds := GetTimeslotDatetimes(ts, now)
		if its.ID < ts.ID && itsStarts.After(tsStarts) {
			return nil, huma.Error400BadRequest("input timeslot can't start after a timeslot with a higher ID.")
		}
		if its.ID > ts.ID && itsStarts.Before(tsStarts) {
			return nil, huma.Error400BadRequest("input timeslot can't start before a timeslot with a lower ID.")
		}
		// input timeslot starts or ends during another
		if itsStarts.After(tsStarts) && itsStarts.Before(tsEnds) || itsEnds.After(tsStarts) && itsEnds.Before(tsEnds) {
			return nil, huma.Error400BadRequest("input timeslot overlaps with an existing timeslot")
		}
	}

	err = db.Queries.UpsertTimeslot(ctx, queries.UpsertTimeslotParams{
		ID:       input.Body.ID,
		StartsAt: itsStarts,
	})
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	return nil, nil
}
