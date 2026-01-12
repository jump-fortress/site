package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"slices"
	"strconv"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

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
	url := fmt.Sprintf("/maps/name/%s/zones/typeindex/%s/%d/records/player/%d/%d", mapName, zoneType, zoneIndex, tempusID, classID)

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
		zoneType = "Map"
		zoneIndex = 1
	} else {
		zoneType = "Map"
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
			RunTime:               tempusPR.RunTime,
			Verified:              true,
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
