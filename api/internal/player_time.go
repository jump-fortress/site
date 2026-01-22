package internal

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

func formatTime(run_time float64) string {
	minutes := int(run_time / 60)
	seconds := math.Mod(run_time, 60)

	return fmt.Sprintf("%02d:%06.3f", minutes, seconds)
}

func getTimeWithPlayerResponse(row queries.SelectCompetitionDivisionPRTimesRow) responses.TimeWithPlayer {
	return responses.TimeWithPlayer{
		Player: responses.PlayerPreview{
			ID:                row.PlayerID,
			Role:              row.Role,
			SteamAvatarUrl:    row.SteamAvatarUrl.String,
			TempusID:          row.TempusID.Int64,
			Country:           row.Country.String,
			CountryCode:       row.CountryCode.String,
			DisplayName:       row.DisplayName.String,
			SoldierDivision:   row.SoldierDivision.String,
			DemoDivision:      row.DemoDivision.String,
			MotwTimeslot:      row.MotwTimeslot.Int64,
			PreferredClass:    row.PreferredClass,
			PreferredLauncher: row.PreferredLauncher.String,
			PreferredMap:      row.PreferredMap.String,
			CreatedAt:         row.CreatedAt_2,
		},
		Time: responses.PlayerTime{
			ID:                    row.ID,
			PlayerID:              row.PlayerID,
			CompetitionDivisionID: row.CompetitionDivisionID,
			TempusTimeID:          row.TempusTimeID.Int64,
			RunTime:               row.RunTime,
			Verified:              row.Verified,
			CreatedAt:             row.CreatedAt,
		},
	}
}

func getCompetitionDivisionTimesResponse(cd queries.CompetitionDivision) responses.CompetitionDivisionTimes {
	return responses.CompetitionDivisionTimes{
		ID:    cd.ID,
		Times: []responses.TimeWithPlayer{},
	}
}

func getCompetitionInProgress(ctx context.Context, id int64) (*queries.Competition, error) {
	competition, err := responses.Queries.SelectCompetition(ctx, id)
	if err != nil {
		return nil, huma.Error500InternalServerError("competition not found")
	}

	// check that competition is after starts_at and not completed
	now := time.Now()
	if competition.StartsAt.After(now) || competition.Complete {
		return nil, huma.Error400BadRequest("competition not in progress")
	}

	return &competition, nil
}

// monthly, motw, quest are only division restricted competitions

// monthly -> cd must match
// motw -> cd must match
// bounty -> no division
// quest -> division must match or be below
//
// playoffs -> no division
// jwc -> no division
// archive -> no division
func isDivisionRestricted(competition_type string) bool {
	if competition_type == "Monthly" || competition_type == "Motw" || competition_type == "Quest" {
		return true
	}
	return false
}

func getTempusPlayerTime(tempusID int64, mapName string, zoneType string, zoneIndex int64, classID int64) (*responses.TempusPlayerTimeResult, error) {
	url := fmt.Sprintf("https://tempus2.xyz/api/v0/maps/name/%s/zones/typeindex/%s/%d/records/player/%d/%d", mapName, zoneType, zoneIndex, tempusID, classID)

	response, err := retryablehttp.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	tempusResponsePlayerTime := &responses.TempusPlayerTimeResultReponse{}
	if err := json.Unmarshal(body, &tempusResponsePlayerTime); err != nil {
		return nil, err
	}

	return &tempusResponsePlayerTime.PlayerPR, nil
}

// todo: only implemented for monthly
func HandlePostSubmitPlayerTime(ctx context.Context, input *responses.CompetitionIDInput) (*struct{}, error) {
	principal, ok := GetPrincipal(ctx)
	if !ok {
		return nil, huma.Error401Unauthorized("a session is required")
	}

	competition, err := getCompetitionInProgress(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	isDivRestrictedCompetition := isDivisionRestricted(competition.Type)

	player, err := responses.Queries.SelectPlayer(ctx, principal.SteamID.String())
	if err != nil {
		return nil, huma.Error500InternalServerError("player not found")
	}

	// check for Tempus ID and division
	if !player.TempusID.Valid {
		return nil, huma.Error400BadRequest("no Tempus ID found")
	}

	var playerDivision string
	var tempusClassID int64
	if competition.Class == "Soldier" {
		tempusClassID = 3
		if player.SoldierDivision.Valid && isDivRestrictedCompetition {
			playerDivision = player.SoldierDivision.String
		} else {
			return nil, huma.Error400BadRequest("no soldier division found")
		}
	} else { // demo
		tempusClassID = 4
		if player.DemoDivision.Valid && isDivRestrictedCompetition {
			playerDivision = player.DemoDivision.String
		} else {
			return nil, huma.Error400BadRequest("no demo division found")
		}
	}

	var zoneType string
	var zoneIndex int64
	// todo: bounty may be course or bonus
	if competition.Type == "Bounty" {
		fmt.Println("todo not implemented - submitting course / bonus zones for a bounty")
		zoneType = "map"
		zoneIndex = 1
	} else {
		zoneType = "map"
		zoneIndex = 1
	}

	competitionDivisions, err := responses.Queries.SelectCompetitionDivisions(ctx, competition.ID)
	if err != nil {
		return nil, huma.Error500InternalServerError(fmt.Sprintf("no divisions found for this %s", competition.Type))
	}

	// check that competition division matches player's division
	var mcd queries.CompetitionDivision
	matched := false
	for _, cd := range competitionDivisions {
		// quests may be submitted by lower division players
		if competition.Type == "Quest" {
			pdi := slices.Index(responses.Divisions, playerDivision)
			cdi := slices.Index(responses.Divisions, cd.Division)
			if pdi == -1 || cdi == -1 {
				return nil, huma.Error500InternalServerError("error finding valid quest division")
			}
			if pdi >= cdi {
				mcd = cd
				matched = true
				break
			}

			// monthly, motw
		} else if cd.Division == playerDivision {
			mcd = cd
			matched = true
			break
		}
	}

	if !matched {
		return nil, huma.Error400BadRequest(fmt.Sprintf("no %s division found for this %s", playerDivision, competition.Type))
	}

	tempusPR, err := getTempusPlayerTime(player.TempusID.Int64, mcd.Map, zoneType, zoneIndex, tempusClassID)
	if err != nil {
		return nil, huma.Error503ServiceUnavailable("couldn't verify with Tempus")
	}

	// check that tempus time is after starts_at and before ends_at
	prUnixSeconds := int64(tempusPR.Date)
	prDate := time.Unix(prUnixSeconds, 0)

	if prDate.After(competition.StartsAt) && prDate.Before(competition.EndsAt) {
		if err := responses.Queries.InsertPlayerTime(ctx, queries.InsertPlayerTimeParams{
			PlayerID:              principal.SteamID.String(),
			CompetitionDivisionID: mcd.ID,
			TempusTimeID: sql.NullInt64{
				Int64: tempusPR.ID,
				Valid: true,
			},
			RunTime:   tempusPR.RunTime,
			Verified:  true,
			CreatedAt: prDate,
		}); err != nil {
			// something went wrong submitting time..
			return nil, err
		}
	} else { // tempus pr is not during competition, request unverified submit
		return nil, huma.Error400BadRequest("Tempus PR wasn't during this competition. submit an unverified time")
	}

	return nil, nil
}

func HandlePostSubmitUnverifiedPlayerTime(ctx context.Context, input *responses.CompetitionIDInput) (*struct{}, error) {
	return nil, huma.Error500InternalServerError("not implemented")

	// check for player division and tempus id
	// check that competition division matches player's division
	// check that competition is after starts_at and *before ends_at*

	// check tempus pr after
	// if tempus pr is faster, submit unverified time
}

// moderator

func HandlePostVerifyPlayerTime(ctx context.Context, input *responses.CompetitionIDInput) (*struct{}, error) {
	if err := responses.Queries.VerifyPlayerTime(ctx, input.ID); err != nil {
		return nil, huma.Error400BadRequest(fmt.Sprintf("couldn't verify time with ID %d", input.ID))
	}

	return nil, nil
}

// admin

// note: functions as SubmitPlayerTime with no Tempus ID and Tempus PR check
func HandlePostCreatePlayerTime(ctx context.Context, input *responses.PlayerTimeInput) (*struct{}, error) {
	competition, err := getCompetitionInProgress(ctx, input.CompetitionID)
	if err != nil {
		return nil, err
	}

	isDivRestrictedCompetition := isDivisionRestricted(competition.Type)

	player, err := responses.Queries.SelectPlayer(ctx, input.PlayerID)
	if err != nil {
		return nil, huma.Error500InternalServerError("player not found")
	}

	var playerDivision string
	if competition.Class == "Soldier" {
		if player.SoldierDivision.Valid && isDivRestrictedCompetition {
			playerDivision = player.SoldierDivision.String
		} else {
			return nil, huma.Error400BadRequest("no soldier division found")
		}
	} else { // demo
		if player.DemoDivision.Valid && isDivRestrictedCompetition {
			playerDivision = player.DemoDivision.String
		} else {
			return nil, huma.Error400BadRequest("no demo division found")
		}
	}

	competitionDivisions, err := responses.Queries.SelectCompetitionDivisions(ctx, competition.ID)
	if err != nil {
		return nil, huma.Error500InternalServerError(fmt.Sprintf("no divisions found for this %s", competition.Type))
	}

	// check that competition division matches player's division
	var mcd queries.CompetitionDivision
	matched := false
	for _, cd := range competitionDivisions {
		// quests may be submitted by lower division players
		if competition.Type == "Quest" {
			pdi := slices.Index(responses.Divisions, playerDivision)
			cdi := slices.Index(responses.Divisions, cd.Division)
			if pdi == -1 || cdi == -1 {
				return nil, huma.Error500InternalServerError("error finding valid quest division")
			}
			if pdi >= cdi {
				mcd = cd
				matched = true
				break
			}

			// monthly, motw
		} else if cd.Division == playerDivision {
			mcd = cd
			matched = true
			break
		}
	}

	if !matched {
		return nil, huma.Error400BadRequest(fmt.Sprintf("no %s division found for this %s", playerDivision, competition.Type))
	}

	if err := responses.Queries.InsertPlayerTime(ctx, queries.InsertPlayerTimeParams{
		PlayerID:              input.PlayerID,
		CompetitionDivisionID: mcd.ID,
		RunTime:               input.RunTime,
		Verified:              true,
	}); err != nil {
		// something went wrong submitting time..
		return nil, err
	}

	return nil, nil
}

func HandlePostDeletePlayerTime(ctx context.Context, input *responses.PlayerTimeIDInput) (*struct{}, error) {
	time, err := responses.Queries.DeletePlayerTime(ctx, input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest(fmt.Sprintf("error deleting a time with ID %d", input.ID))
	}

	jsonTime, err := json.Marshal(time)
	if err != nil {
		return nil, err
	}

	if err := responses.Queries.InsertDeletedRecord(ctx, queries.InsertDeletedRecordParams{
		SourceTable: "player_time",
		SourceID:    strconv.FormatInt(time.ID, 10),
		Data:        json.RawMessage(jsonTime),
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

func HandleGetCompetitionPlayerTimes(ctx context.Context, input *responses.CompetitionTypeAndIDInput) (*responses.CompetitionTimesOutput, error) {
	cid, err := getCompetitionID(ctx, input.CompetitionType, input.ID)
	if err != nil {
		return nil, err
	}

	//c, err := responses.Queries.SelectCompetition(ctx, cid)
	// if err != nil {
	// 	return nil, err
	// }

	competitionDivisions, err := responses.Queries.SelectCompetitionDivisions(ctx, cid)
	if err != nil {
		return nil, huma.Error400BadRequest(fmt.Sprintf("no competition found with id %d", cid))
	}

	resp := &responses.CompetitionTimesOutput{
		Body: []responses.CompetitionDivisionTimes{},
	}

	for _, cd := range competitionDivisions {
		cdtResponse := getCompetitionDivisionTimesResponse(cd)

		times, err := responses.Queries.SelectCompetitionDivisionPRTimes(ctx, cd.ID)
		if err != nil {
			return nil, err
		}

		for _, t := range times {
			cdtResponse.Times = append(cdtResponse.Times, getTimeWithPlayerResponse(t))
		}

		resp.Body = append(resp.Body, cdtResponse)
	}

	return resp, nil
}

// dev
func HandlePostUpdatePlayerTimesTempus(ctx context.Context, input *responses.CompetitionTypeAndIDInput) (*struct{}, error) {
	cid, err := getCompetitionID(ctx, input.CompetitionType, input.ID)
	if err != nil {
		return nil, err
	}

	c, err := responses.Queries.SelectCompetition(ctx, cid)
	if err != nil {
		return nil, err
	}

	competitionDivisions, err := responses.Queries.SelectCompetitionDivisions(ctx, cid)
	if err != nil {
		return nil, huma.Error400BadRequest(fmt.Sprintf("no competition found with id %d", cid))
	}

	for _, cd := range competitionDivisions {

		times, err := responses.Queries.SelectCompetitionDivisionPRTimes(ctx, cd.ID)
		if err != nil {
			return nil, err
		}

		for _, t := range times {

			tt, err := getTempusPlayerTime(t.TempusID.Int64, cd.Map, "map", 1, 3)
			if err != nil {
				return nil, err
			}

			time.Sleep(time.Second)

			// check that tempus time is after starts_at and before ends_at
			prUnixSeconds := int64(tt.Date)
			prDate := time.Unix(prUnixSeconds, 0)
			fmt.Printf("%s: %f", t.DisplayName.String, tt.RunTime)

			if prDate.After(c.StartsAt) && prDate.Before(c.EndsAt.AddDate(0, 0, 1)) {
				fmt.Println("valid")
				err := responses.Queries.UpdatePlayerTimeFromTempus(ctx, queries.UpdatePlayerTimeFromTempusParams{
					RunTime:   tt.RunTime,
					CreatedAt: prDate,
					TempusTimeID: sql.NullInt64{
						Int64: int64(tt.RunTime),
						Valid: true,
					},
					ID: t.ID,
				})
				if err != nil {
					return nil, err
				}
			}

		}
	}

	return nil, nil
}
