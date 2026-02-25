package routes

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jump-fortress/site/db"
	"github.com/jump-fortress/site/db/queries"
	"github.com/jump-fortress/site/internal/principal"
	"github.com/jump-fortress/site/internal/rows"
	"github.com/jump-fortress/site/internal/tempus"
	"github.com/jump-fortress/site/models"
)

var (
	timeSubmitExtension = time.Hour * 24     // player submissions grace period
	timeManageExtension = time.Hour * 24 * 7 // mod / admin updates grace period
)

// validate that a time isn't already submitted, and isn't larger duration than current PR
func ValidateTimeExistsAndPR(ctx context.Context, leaderboard_id int64, player_id string, duration float64) error {
	timeExists, err := db.Queries.SelectTimeExists(ctx, queries.SelectTimeExistsParams{
		LeaderboardID: leaderboard_id,
		PlayerID:      player_id,
		Duration:      duration,
	})
	if err != nil {
		return models.WrapDBErr(err)
	}
	if timeExists == 1 {
		return huma.Error400BadRequest("time has already been submitted")
	}

	timesSubmitted, err := db.Queries.CountPlayerTimesFromLeaderboard(ctx, queries.CountPlayerTimesFromLeaderboardParams{
		LeaderboardID: leaderboard_id,
		PlayerID:      player_id,
	})

	// get time
	if timesSubmitted != 0 {
		twps, err := db.Queries.SelectPRTimesFromLeaderboard(ctx, leaderboard_id)
		if err != nil {
			return models.WrapDBErr(err)
		}

		for _, twp := range twps {
			if twp.Player.ID == player_id {
				if twp.Time.Duration < duration {
					return huma.Error400BadRequest(fmt.Sprintf("submitted time (%.3f seconds) is slower than your current PR (%.3f seconds)", duration, twp.Time.Duration))
				}
				return nil
			}
		}
		return huma.Error404NotFound(fmt.Sprintf("%d time(s) found submitted, but couldn't find their details.", timesSubmitted))
	}
	return nil
}

// get a player, leaderboard, and event
func GetEventDetailsForLeaderboard(ctx context.Context, leaderboard_id int64, player_id string) (*queries.Player, *queries.Leaderboard, *queries.Event, error) {
	player, err := db.Queries.SelectPlayer(ctx, player_id)
	if err != nil {
		return nil, nil, nil, models.WrapDBErr(err)
	}
	leaderboard, err := db.Queries.SelectLeaderboard(ctx, leaderboard_id)
	if err != nil {
		return nil, nil, nil, models.WrapDBErr(err)
	}
	event, err := db.Queries.SelectEventFromLeaderboardID(ctx, leaderboard.ID)
	if err != nil {
		return nil, nil, nil, models.WrapDBErr(err)
	}

	// Tempus PRs have a timestamp, so checking if a competition has ended isn't necessary
	now := time.Now().UTC()
	if event.StartsAt.After(now) {
		return nil, nil, nil, huma.Error400BadRequest(fmt.Sprintf("%s hasn't started", event.Kind))
	}

	var playerDiv sql.NullString
	if event.Class == "Soldier" {
		playerDiv = player.SoldierDiv
	} else {
		playerDiv = player.DemoDiv
	}

	// validate player div if leaderboard has div
	if leaderboard.Div.Valid {
		if !playerDiv.Valid {
			return nil, nil, nil, huma.Error400BadRequest(fmt.Sprintf("missing a %s div", event.Class))
		}
		if leaderboard.Div.String != playerDiv.String {
			return nil, nil, nil, huma.Error400BadRequest(fmt.Sprintf("can't submit a %s time to a %s leaderboard", playerDiv.String, leaderboard.Div.String))
		}
	}

	return &player, &leaderboard, &event, nil
}

func HandleSubmitTime(ctx context.Context, input *models.LeaderboardIDInput) (*struct{}, error) {
	principal, ok := principal.Get(ctx)
	if !ok {
		return nil, models.SessionErr()
	}
	player, leaderboard, event, err := GetEventDetailsForLeaderboard(ctx, input.ID, principal.SteamID.String())
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	if event.EndsAt.Add(timeSubmitExtension).Before(now) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("%s has ended. please contact a mod if you have an old time to submit!", event.Kind))
	}

	// validate player Tempus ID
	if !player.TempusID.Valid {
		return nil, huma.Error400BadRequest("missing a Tempus ID")
	}
	var tempusClassID int64
	if event.Class == "Soldier" {
		tempusClassID = 3
	} else {
		tempusClassID = 4
	}

	// todo: support for multiple zone types found in other event kinds (course, bonus)
	tt, err := tempus.GetPR(leaderboard.Map, player.TempusID.Int64, tempusClassID)
	if err != nil {
		return nil, models.WrapTempusErr(err)
	}

	// validate Tempus PR is during event
	ttDate := time.Unix(int64(tt.Date), 0)
	if ttDate.Before(event.StartsAt) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("Tempus PR was before this event started! (%s) please submit an unverified time if you're submitting a non-PR time", ttDate.String()))
	}
	// for motws, additionally check if Tempus PR is during player's timeslot
	if event.Kind == "motw" {
		playerTimeslot, err := db.Queries.SelectPlayerTimeslot(ctx, principal.SteamID.String())
		if err != nil {
			return nil, huma.Error400BadRequest("no timeslot found, please select one when a motw isn't in progress")
		}

		eventPts := GetTimeslotDatetimes(playerTimeslot.MotwTimeslot, event.StartsAt)
		if ttDate.Before(eventPts.StartsAt) || ttDate.After(eventPts.EndsAt) {
			return nil, huma.Error400BadRequest(fmt.Sprintf("Tempus PR wasn't during your timeslot! (%s) please submit an unverified time if submitting a non-PR time.", ttDate.String()))
		}
	}

	err = ValidateTimeExistsAndPR(ctx, leaderboard.ID, player.ID, tt.Duration)
	if err != nil {
		return nil, err
	}

	// submit verified time!
	err = db.Queries.InsertTime(ctx, queries.InsertTimeParams{
		LeaderboardID: leaderboard.ID,
		PlayerID:      player.ID,
		TempusTimeID: sql.NullInt64{
			Int64: tt.ID,
			Valid: true,
		},
		Duration: tt.Duration,
		Verified: true,
	})
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	return nil, nil
}

func HandleSubmitUnverifiedTime(ctx context.Context, input *models.UnverifiedTimeInput) (*struct{}, error) {
	principal, ok := principal.Get(ctx)
	if !ok {
		return nil, models.SessionErr()
	}
	player, leaderboard, event, err := GetEventDetailsForLeaderboard(ctx, input.ID, principal.SteamID.String())
	if err != nil {
		return nil, err
	}

	// check input is correct format
	var minutes int64
	var seconds float64

	parsed, err := fmt.Sscanf(input.RunTime, "%d:%f", &minutes, &seconds)
	if parsed != 2 {
		return nil, huma.Error400BadRequest("couldn't parse your time correctly. format: MM:SS.ss (include 0 for minutes if time is less than 60 seconds)")
	}

	var duration float64 = float64(minutes*60) + seconds

	// additionally don't allow submitting before a player's timeslot
	if event.Kind == "motw" {
		playerTimeslot, err := db.Queries.SelectPlayerTimeslot(ctx, principal.SteamID.String())
		if err != nil {
			return nil, huma.Error400BadRequest("no timeslot found, please select one when a motw isn't in progress")
		}

		now := time.Now().UTC()
		eventPts := GetTimeslotDatetimes(playerTimeslot.MotwTimeslot, event.StartsAt)
		if eventPts.StartsAt.After(now) {
			return nil, huma.Error400BadRequest("motw hasn't started in your timeslot yet")
		}
		if eventPts.EndsAt.Before(now) {
			return nil, huma.Error400BadRequest("motw has ended in your timeslot. please contact a mod if you have an old time to submit!")
		}
	}

	now := time.Now().UTC()
	if event.EndsAt.Add(timeSubmitExtension).Before(now) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("%s has ended. please contact a mod if you have an old time to submit!", event.Kind))
	}

	// validate player Tempus ID
	if !player.TempusID.Valid {
		return nil, huma.Error400BadRequest("missing a Tempus ID")
	}
	var tempusClassID int64
	if event.Class == "Soldier" {
		tempusClassID = 3
	} else {
		tempusClassID = 4
	}

	// todo: support for multiple zone types found in other event kinds (course, bonus)
	tt, err := tempus.GetPR(leaderboard.Map, player.TempusID.Int64, tempusClassID)
	if err != nil {
		return nil, models.WrapTempusErr(err)
	}

	// submitted time is faster than Tempus PR
	if tt.Duration > duration {
		return nil, huma.Error400BadRequest(fmt.Sprintf("couldn't submit a time faster than your tempus PR (%.3f seconds)", tt.Duration))
	}

	err = ValidateTimeExistsAndPR(ctx, leaderboard.ID, player.ID, duration)
	if err != nil {
		return nil, err
	}

	// validate input duration is positive and less than 10 hours
	if duration < 0 || duration > 60*60*10 {
		return nil, models.InvalidDurationErr(duration)
	}

	err = db.Queries.InsertTime(ctx, queries.InsertTimeParams{
		LeaderboardID: leaderboard.ID,
		PlayerID:      player.ID,
		TempusTimeID: sql.NullInt64{
			Int64: 0,
			Valid: false,
		},
		Duration: duration,
		Verified: false,
	})

	return nil, nil
}

func HandleGetEventPR(ctx context.Context, input *models.EventIDInput) (*models.TimeWithPlayerOutput, error) {
	principal, ok := principal.Get(ctx)
	if !ok {
		return nil, models.SessionErr()
	}

	leaderboards, err := db.Queries.SelectLeaderboards(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	for _, l := range leaderboards {
		twps, err := db.Queries.SelectPRTimesFromLeaderboard(ctx, l.ID)
		if err != nil {
			return nil, models.WrapDBErr(err)
		}

		for _, twp := range twps {
			if twp.Player.ID == principal.SteamID.String() {
				resp := &models.TimeWithPlayerOutput{
					Body: models.TimeWithPlayer{
						Time:     models.GetTimeResponse(twp.Time),
						Player:   models.GetPlayerResponse(twp.Player, false),
						Position: twp.TimePosition,
					},
				}
				return resp, nil
			}
		}
	}
	return nil, huma.Error404NotFound("no PR found")
}

func HandleGetPlayerPRs(ctx context.Context, input *models.PlayerIDInput) (*models.EventLeaderboardTimesOutput, error) {
	elts, err := db.Queries.SelectParticipatedEvents(ctx, input.PlayerID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	resp := &models.EventLeaderboardTimesOutput{
		Body: []models.EventLeaderboardTime{},
	}

	for _, elt := range elts {
		// hide motw times if motw hasn't ended
		sensitive := motwNotEnded(elt.Event.Kind, elt.Event.EndsAt)
		if sensitive {
			continue
		}
		// get time rank
		twps, err := db.Queries.SelectPRTimesFromLeaderboard(ctx, elt.Leaderboard.ID)
		if err != nil {
			return nil, models.WrapDBErr(err)
		}

		for _, twp := range twps {
			if twp.Player.ID == input.PlayerID {
				resp.Body = append(resp.Body, models.EventLeaderboardTime{
					Event:       models.GetEventResponse(elt.Event),
					Leaderboard: models.GetLeaderboardResponse(elt.Leaderboard, false),
					Time:        models.GetTimeResponse(elt.Time),
					Position:    twp.TimePosition,
				})
			}
		}
	}

	return resp, nil
}

// mod

func HandleSubmitPlayerTime(ctx context.Context, input *models.PlayerTimeInput) (*struct{}, error) {
	// skips Tempus checks
	// times submitted by a mod are verified

	player, leaderboard, event, err := GetEventDetailsForLeaderboard(ctx, input.LeaderboardID, input.PlayerID)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	if event.EndsAt.Before(now.Add(timeManageExtension)) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("%s ended more than one week ago. can't submit a time this old", event.Kind))
	}

	// validate input duration is positive and less than 10 hours
	if input.Duration < 0 || input.Duration > 60*60*10 {
		return nil, models.InvalidDurationErr(input.Duration)
	}

	err = ValidateTimeExistsAndPR(ctx, leaderboard.ID, player.ID, input.Duration)
	if err != nil {
		return nil, err
	}

	err = db.Queries.InsertTime(ctx, queries.InsertTimeParams{
		LeaderboardID: leaderboard.ID,
		PlayerID:      player.ID,
		TempusTimeID: sql.NullInt64{
			Int64: 0,
			Valid: false,
		},
		Duration: input.Duration,
		Verified: true,
	})

	return nil, nil
}

func HandleVerifyPlayerTime(ctx context.Context, input *models.TimeIDInput) (*struct{}, error) {
	err := db.Queries.VerifyTime(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}
	return nil, nil
}

func HandleDeletePlayerTime(ctx context.Context, input *models.TimeIDInput) (*struct{}, error) {
	time, err := db.Queries.SelectTime(ctx, input.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	if time.TempusTimeID.Valid {
		return nil, huma.Error400BadRequest("time has a valid Tempus PR attached. can't remove this time")
	}

	err = db.Queries.DeleteTime(ctx, time.ID)
	if err != nil {
		return nil, models.WrapDBErr(err)
	}

	err = rows.InsertDeleted(ctx, time, "time", time.ID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
