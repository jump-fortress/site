package routes

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jump-fortress/site/db"
	"github.com/jump-fortress/site/db/queries"
	"github.com/jump-fortress/site/internal/rows"
	"github.com/jump-fortress/site/models"
)

func getEndsAt(kind string, starts_at time.Time) time.Time {
	switch kind {
	case "monthly":
		return starts_at.Add(time.Hour * 48)
	default:
		// should never be reached
		return starts_at.Add(time.Hour * 48)
	}
}

func HandleGetEvent(ctx context.Context, input *models.EventKindAndIDInput) (*models.EventWithLeaderboardsOutput, error) {
	// get event and leaderboards
	els, err := db.Queries.SelectEventLeaderboards(ctx, queries.SelectEventLeaderboardsParams{
		Kind:   input.Kind,
		KindID: input.KindID,
	})
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	now := time.Now()
	// check for player role, and only include upcoming event details for admin role
	admin, err := ValidateAdminRole(ctx)
	if err != nil {
		return nil, err
	}
	if !admin && els[0].Event.VisibleAt.After(now) {
		return nil, huma.Error400BadRequest("event not visible")
	}

	resp := &models.EventWithLeaderboardsOutput{
		Body: models.GetEventWithLeaderboardsResponse(els, !admin),
	}

	return resp, nil

}

func HandleGetEventKinds(ctx context.Context, input *models.EventKindInput) (*models.EventsWithLeaderboardsOutput, error) {
	// validate input
	if !slices.Contains(models.EventKinds, input.Kind) {
		return nil, models.EventKindErr(input.Kind)
	}

	// check for player role, and only include upcoming event details for admin role
	admin, err := ValidateAdminRole(ctx)
	if err != nil {
		return nil, err
	}

	// get event kinds
	events, err := db.Queries.SelectEventKinds(ctx, input.Kind)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	now := time.Now()
	resp := &models.EventsWithLeaderboardsOutput{
		Body: []models.EventWithLeaderboards{},
	}

	for _, e := range events {
		// skip non-visible events
		if !admin && e.VisibleAt.After(now) {
			continue
		}

		// append leaderboards to each event
		ls, err := db.Queries.SelectLeaderboards(ctx, e.ID)
		if err != nil {
			return nil, models.WrapDBErr(err)
		}
		els := []queries.SelectEventLeaderboardsRow{}
		for _, l := range ls {
			els = append(els, queries.SelectEventLeaderboardsRow{
				Event:       e,
				Leaderboard: l,
			})
		}
		resp.Body = append(resp.Body, models.GetEventWithLeaderboardsResponse(els, !admin))
	}

	if len(resp.Body) == 0 {
		return nil, huma.Error400BadRequest("no events visible")
	}
	return resp, nil
}

// admin

func HandleCreateEvent(ctx context.Context, input *models.EventInput) (*struct{}, error) {
	// validate event info
	ie := input.Body
	if !slices.Contains(models.EventKinds, ie.Kind) {
		return nil, models.EventKindErr(ie.Kind)
	}
	if ie.PlayerClass != "Soldier" && ie.PlayerClass != "Demo" {
		return nil, models.PlayerClassErr(ie.PlayerClass)
	}

	// time checks
	now := time.Now()
	if ie.EndsAt.Before(ie.StartsAt) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("event can't end before it starts (%s)", ie.EndsAt.String()))
	}
	if ie.StartsAt.Before(now) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("event can't start in the past (%s)", ie.StartsAt.String()))
	}
	if ie.StartsAt.Before(ie.VisibleAt) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("event can't start before it's visible (%s)", ie.VisibleAt.String()))
	}

	// set ID for next event
	kindID, err := db.Queries.CountEventKinds(ctx, ie.Kind)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}
	kindID++

	endsAt := getEndsAt(ie.Kind, ie.StartsAt)

	// create
	err = db.Queries.InsertEvent(ctx, queries.InsertEventParams{
		Kind:      ie.Kind,
		KindID:    kindID,
		Class:     ie.PlayerClass,
		VisibleAt: ie.VisibleAt,
		StartsAt:  ie.StartsAt,
		EndsAt:    endsAt,
	})
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	return nil, nil
}

func HandleUpdateEvent(ctx context.Context, input *models.EventInput) (*struct{}, error) {
	// validate event info
	ie := input.Body
	event, err := db.Queries.SelectEvent(ctx, ie.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	if ie.Kind != event.Kind || ie.KindID != event.KindID {
		return nil, huma.Error400BadRequest("can't modify the event kind or kind_id")
	}

	now := time.Now()
	if event.StartsAt.Before(now) {
		return nil, huma.Error400BadRequest("event has already started")
	}
	if ie.EndsAt.Before(ie.StartsAt) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("event can't end before it starts (%s)", ie.EndsAt.String()))
	}
	if ie.StartsAt.Before(now) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("event can't start in the past (%s)", ie.StartsAt.String()))
	}
	if ie.StartsAt.Before(ie.VisibleAt) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("event can't start before it's visible (%s)", ie.VisibleAt.String()))
	}

	return nil, nil
}

func HandleCancelEvent(ctx context.Context, input *models.EventIDInput) (*struct{}, error) {
	// check that the event hasn't started
	event, err := db.Queries.SelectEvent(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	now := time.Now()
	if event.StartsAt.Before(now) {
		return nil, huma.Error400BadRequest("event has already started")
	}

	err = db.Queries.DeleteEvent(ctx, event.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	// backup event
	return nil, rows.InsertDeleted(ctx, event, "event", event.ID)
}
