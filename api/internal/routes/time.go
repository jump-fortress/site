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
		return huma.Error400BadRequest(fmt.Sprintf("time has already been submitted (%f.3 seconds)", duration))
	}

	timesSubmitted, err := db.Queries.CountPlayerTimesFromLeaderboard(ctx, queries.CountPlayerTimesFromLeaderboardParams{
		LeaderboardID: leaderboard_id,
		PlayerID:      player_id,
	})
	if timesSubmitted != 0 {
		prDuration, err := db.Queries.SelectPRTime(ctx, queries.SelectPRTimeParams{
			LeaderboardID: leaderboard_id,
			PlayerID:      player_id,
		})
		if err != nil {
			return models.WrapDBErr(err)
		}
		if prDuration < duration {
			return huma.Error400BadRequest(fmt.Sprintf("submitted time (%f.3 seconds) is slower than your current PR (%f.3 seconds)", duration, prDuration))
		}
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
	now := time.Now()
	if event.StartsAt.After(now) {
		return nil, nil, nil, huma.Error400BadRequest(fmt.Sprintf("%s hasn't started", event.Kind))
	}

	var playerDiv sql.NullString
	if event.Class == "soldier" {
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

	// validate player Tempus ID
	if !player.TempusID.Valid {
		return nil, huma.Error400BadRequest("missing a Tempus ID")
	}
	var tempusClassID int64
	if event.Class == "soldier" {
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
	if ttDate.After(event.EndsAt) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("Tempus PR was after this event ended! (%s)", ttDate.String()))
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
	// an unverified time skips a few checks from verified ones..
	// 1. a Tempus ID is not required
	// 2. a Tempus class isn't needed, and no request is made to Tempus
	// - a request still could be made to Tempus if a Tempus ID is provided, to make sure there's no valid Tempus PR faster than the submitted time
	// - but this should be rare, and a result of user error. less external API requests are good!

	principal, ok := principal.Get(ctx)
	if !ok {
		return nil, models.SessionErr()
	}
	player, leaderboard, event, err := GetEventDetailsForLeaderboard(ctx, input.ID, principal.SteamID.String())
	if err != nil {
		return nil, err
	}

	// one day "grace period" for submitting an unverified time
	now := time.Now()
	if event.EndsAt.Before(now.Add(time.Hour * 24)) {
		return nil, huma.Error400BadRequest(fmt.Sprintf("%s has ended. please contact a mod if you have an old time to submit!", event.Kind))
	}
	err = ValidateTimeExistsAndPR(ctx, leaderboard.ID, player.ID, input.Duration)
	if err != nil {
		return nil, err
	}

	// validate input duration is positive and less than 10 hours
	if input.Duration < 0 || input.Duration > 60*60*10 {
		return nil, models.InvalidDurationErr(input.Duration)
	}

	err = db.Queries.InsertTime(ctx, queries.InsertTimeParams{
		LeaderboardID: leaderboard.ID,
		PlayerID:      player.ID,
		TempusTimeID: sql.NullInt64{
			Int64: 0,
			Valid: false,
		},
		Duration: input.Duration,
		Verified: false,
	})

	return nil, nil
}

// mod

func HandleSubmitPlayerTime(ctx context.Context, input *models.PlayerTimeInput) (*struct{}, error) {
	// skips Tempus checks
	// times submitted by a mod are verified

	player, leaderboard, event, err := GetEventDetailsForLeaderboard(ctx, input.LeaderboardID, input.PlayerID)
	if err != nil {
		return nil, err
	}

	// one week "grace period" for submitting a verified time
	now := time.Now()
	if event.EndsAt.Before(now.Add(time.Hour * 24 * 7)) {
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
