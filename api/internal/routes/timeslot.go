package routes

import (
	"context"
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

func GetTimeslotDatetimes(timeslot queries.MotwTimeslot) (time.Time, time.Time) {
	now := time.Now()
	starts := time.Date(now.Year(), now.Month(), now.Day(), timeslot.StartsAt.Hour(), timeslot.StartsAt.Minute(), timeslot.StartsAt.Second(), timeslot.StartsAt.Nanosecond(), time.UTC)
	ends := starts.Add(MotwDuration)
	return starts, ends
}

func HandleUpdateTimeslotPref(ctx context.Context, input *models.TimeslotIDInput) (*struct{}, error) {
	principal, ok := principal.Get(ctx)
	if !ok {
		return nil, models.SessionErr()
	}

	err := db.Queries.UpdatePlayerTimeslot(ctx, queries.UpdatePlayerTimeslotParams{
		TimeslotID: input.ID,
		PlayerID:   principal.SteamID.String(),
	})
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	return nil, nil
}

// admin

func HandleUpdateTimeslot(ctx context.Context, input *models.TimeslotInput) (*struct{}, error) {
	its := input.Body
	itsStarts, itsEnds := GetTimeslotDatetimes(queries.MotwTimeslot(its))

	// validate input timeslot doesn't overlap with other timeslots
	timeslots, err := db.Queries.SelectTimeslots(ctx)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	for _, ts := range timeslots {
		// skip checking same ID
		if ts.ID == input.Body.ID {
			continue
		}
		tsStarts, tsEnds := GetTimeslotDatetimes(ts)
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
