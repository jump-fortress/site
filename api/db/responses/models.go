package responses

import (
	"time"
)

// types matching their database name (ex. "Player") are mapped database types
// "Preview" suffixed types are database types missing sensitive or unneccessary properties
// "Input" suffixed types are used for url parameters
// "Output" suffixed typed are used for api responses

var Divisions = []string{"Diamond", "Platinum", "Gold", "Silver", "Bronze", "Steel", "Wood"}

type PlayerIDInput struct {
	ID string `path:"id" minimum:"1" doc:"player ID, SteamID64"`
}

type PlayerPreview struct {
	ID                string    `json:"id"`
	Role              string    `json:"role"`
	SteamAvatarUrl    string    `json:"steam_avatar_url"`
	TempusID          int64     `json:"tempus_id,omitempty"`
	Country           string    `json:"country,omitempty"`
	CountryCode       string    `json:"country_code,omitempty"`
	DisplayName       string    `json:"display_name"`
	SoldierDivision   string    `json:"soldier_division,omitempty"`
	DemoDivision      string    `json:"demo_division,omitempty"`
	PreferredClass    string    `json:"preferred_class"`
	PreferredLauncher string    `json:"preferred_launcher"`
	PreferredMap      string    `json:"preferred_map,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
}

type PlayerRequestPreview struct {
	RequestType   string    `json:"request_type"`
	RequestString string    `json:"request_string"`
	CreatedAt     time.Time `json:"created_at"`
}

type Player struct {
	ID                string    `json:"id"`
	Role              string    `json:"role"`
	SteamAvatarUrl    string    `json:"steam_avatar_url"`
	SteamTradeToken   string    `json:"steam_trade_token,omitempty"`
	TempusID          int64     `json:"tempus_id,omitempty"`
	Country           string    `json:"country,omitempty"`
	CountryCode       string    `json:"country_code,omitempty"`
	DiscordID         string    `json:"discord_id,omitempty"`
	DisplayName       string    `json:"display_name"`
	SoldierDivision   string    `json:"soldier_division,omitempty"`
	DemoDivision      string    `json:"demo_division,omitempty"`
	PreferredClass    string    `json:"preferred_class"`
	PreferredLauncher string    `json:"preferred_launcher"`
	PreferredMap      string    `json:"preferred_map,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
}

type PlayerRequest struct {
	ID            int64     `json:"id"`
	PlayerID      string    `json:"player_id"`
	RequestType   string    `json:"request_type"`
	RequestString string    `json:"request_string,omitempty"`
	Pending       bool      `json:"pending"`
	CreatedAt     time.Time `json:"created_at"`
}

type PlayerWithRequest struct {
	Player  Player        `json:"player"`
	Request PlayerRequest `json:"request"`
}

type PlayerPreviewOutput struct {
	Body PlayerPreview
}

type PlayerPreviewsOutput struct {
	Body []PlayerPreview
}

type PlayerOutput struct {
	Body Player
}

type PlayersOutput struct {
	Body []Player
}

type PlayerRequestPreviewOutput struct {
	Body PlayerRequestPreview
}

type PlayerRequestPreviewsOutput struct {
	Body []PlayerRequestPreview
}

type PlayerRequestsOutput struct {
	Body []PlayerRequest
}

type PlayersWithRequestOutput struct {
	Body []PlayerWithRequest
}

type PlayerPoints struct {
	SoldierPoints PlayerClassPoints `json:"soldier"`
	DemoPoints    PlayerClassPoints `json:"demo"`
}

type PlayerClassPoints struct {
	Total        int64 `json:"total"`
	Last3Monthly int64 `json:"last_3_monthly"`
	Last9Motw    int64 `json:"last_9_motw"`
}

// todo: a separate type to return competition placements
type PlayerProfile struct {
	PlayerPreview PlayerPreview `json:"player"`
	PlayerPoints  PlayerPoints  `json:"points"`
}

type PlayerProfileOutput struct {
	Body PlayerProfile
}

type Session struct {
	ID             string `json:"id"`
	Role           string `json:"role"`
	DisplayName    string `json:"displayName"`
	SteamAvatarURL string `json:"steamAvatarUrl"`
}

type SessionOutput struct {
	Body Session
}

type ClassNameInput struct {
	Class string `path:"class" enum:"Soldier,Demo"`
}

type LauncherNameInput struct {
	Launcher string `path:"launcher" enum:"Stock,Original,Mangler,None"`
}

type MapNameInput struct {
	Map string `path:"map"`
}

type DisplayNameInput struct {
	ID   string `path:"id" minimum:"1" doc:"player ID, SteamID64"`
	Name string `path:"name" doc:"new display name"`
}

type TempusIDInput struct {
	TempusID int64 `path:"tempus_id" minimum:"1" doc:"Tempus ID"`
}

type TempusPlayerInfo struct {
	TempusName  string `json:"name"`
	TempusID    int64  `json:"id"`
	SteamID     string `json:"steamid"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

// used to read in player_info property from the Tempus API
type TempusPlayerInfoResponse struct {
	PlayerInfo TempusPlayerInfo `json:"player_info"`
}

type TempusResponseMaps struct {
	Maps []TempusMapInfo
}

// note: TempusTier and TempusRating have the same properties
// they still mean different things so it may be good to have them separate
type TempusMapInfo struct {
	ID     int64        `json:"id"`
	Name   string       `json:"name"`
	Zones  TempusZones  `json:"zone_counts"`
	Tier   TempusTier   `json:"tier_info"`
	Rating TempusRating `json:"rating_info"`
}

type TempusZones struct {
	Bonus  int64 `json:"bonus"`
	Course int64 `json:"course"`
}

type TempusTier struct {
	Soldier int64 `json:"3"`
	Demo    int64 `json:"4"`
}

type TempusRating struct {
	Soldier int64 `json:"3"`
	Demo    int64 `json:"4"`
}

type Map struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Courses       int64  `json:"courses,omitempty"`
	Bonuses       int64  `json:"bonuses,omitempty"`
	SoldierTier   int64  `json:"soldier_tier"`
	DemoTier      int64  `json:"demo_tier"`
	SoldierRating int64  `json:"soldier_rating"`
	DemoRating    int64  `json:"demo_rating"`
}

type MapsOutput struct {
	Body []Map
}

type MapNamesOutput struct {
	Body []string
}

type SteamTradeURL struct {
	Url string `path:"url"`
}

type PlayerRequestInput struct {
	RequestType   string `path:"request_type" enum:"Display Name Change,Soldier Placement,Demo Placement"`
	RequestString string `path:"request_string" required:"false"`
}

type PlayerRequestIDInput struct {
	ID int64 `path:"id" doc:"request ID"`
}

type UpdateDivisionInput struct {
	ID       string `path:"id" minimum:"1" doc:"player ID, SteamID64"`
	Division string `path:"division"`
}

type CompetitionInput struct {
	Class    string    `json:"class"`
	StartsAt time.Time `json:"starts_at"`
}

type CompetitionDivisionInput struct {
	Division string `json:"division"`
	Map      string `json:"map"`
}

type Competition struct {
	ID        int64     `json:"id"`
	Class     string    `json:"class"`
	StartsAt  time.Time `json:"starts_at"`
	EndsAt    time.Time `json:"ends_at"`
	CreatedAt time.Time `json:"created_at"`
}

type CompetitionDivision struct {
	ID            int64  `json:"id"`
	CompetitionID int64  `json:"competition_id"`
	Division      string `json:"division"`
	Map           string `json:"map"`
}

type MonthlyInput struct {
	Body struct {
		Competition CompetitionInput           `json:"competition"`
		Divisions   []CompetitionDivisionInput `json:"divisions"`
	}
}

type Monthly struct {
	Competition Competition           `json:"competition"`
	Divisions   []CompetitionDivision `json:"divisions"`
}

type MonthliesOutput struct {
	Body []Monthly
}
