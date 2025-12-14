package responses

import (
	"time"
)

var Divisions = []string{"Diamond", "Platinum", "Gold", "Silver", "Bronze", "Steel", "Wood"}

type PlayerIDInput struct {
	ID string `path:"id" minimum:"1" doc:"player ID, SteamID64"`
}

type Player struct {
	ID                string    `json:"id"`
	Role              string    `json:"role"`
	SteamAvatarUrl    string    `json:"steam_avatar_url"`
	TempusID          int64     `json:"tempus_id,omitempty"`
	Country           string    `json:"country"`
	CountryCode       string    `json:"country_code"`
	DisplayName       string    `json:"display_name"`
	SoldierDivision   string    `json:"soldier_division,omitempty"`
	DemoDivision      string    `json:"demo_division,omitempty"`
	PreferredClass    string    `json:"preferred_class"`
	PreferredLauncher string    `json:"preferred_launcher"`
	CreatedAt         time.Time `json:"created_at"`
}

type FullPlayer struct {
	ID                string    `json:"id"`
	Role              string    `json:"role"`
	SteamAvatarUrl    string    `json:"steam_avatar_url"`
	SteamTradeToken   string    `json:"steam_trade_token,omitempty"`
	TempusID          int64     `json:"tempus_id,omitempty"`
	Country           string    `json:"country"`
	CountryCode       string    `json:"country_code"`
	DiscordID         string    `json:"discord_id,omitempty"`
	DisplayName       string    `json:"display_name"`
	SoldierDivision   string    `json:"soldier_division,omitempty"`
	DemoDivision      string    `json:"demo_division,omitempty"`
	PreferredClass    string    `json:"preferred_class"`
	PreferredLauncher string    `json:"preferred_launcher"`
	CreatedAt         time.Time `json:"created_at"`
}

type SelfPlayerRequest struct {
	RequestType   string    `json:"request_type"`
	RequestString string    `json:"request_string"`
	CreatedAt     time.Time `json:"created_at"`
}

type PlayerOutput struct {
	Body Player
}

type ManyPlayersOutput struct {
	Body []Player
}

type FullPlayerOutput struct {
	Body FullPlayer
}

type ManyFullPlayersOutput struct {
	Body []FullPlayer
}

type SelfPlayerRequestOutput struct {
	Body SelfPlayerRequest
}

type ManySelfPlayerRequestsOutput struct {
	Body []SelfPlayerRequest
}

type PlayerPoints struct {
	Total        int64 `json:"total"`
	Last3Monthly int64 `json:"last_3_monthly"`
	Last9Motw    int64 `json:"last_9_motw"`
}

type PlayerProfile struct {
	Player
	SoldierPoints PlayerPoints `json:"soldier_points"`
	DemoPoints    PlayerPoints `json:"demo_points"`
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

// todo: regex for Name
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

type TempusResponsePlayerInfo struct {
	PlayerInfo TempusPlayerInfo `json:"player_info"`
}

type SteamTradeURL struct {
	Url string `path:"url"`
}

type PlayerRequestInput struct {
	RequestType   string `path:"request_type" enum:"Display Name Change,Soldier Placement,Demo Placement"`
	RequestString string `path:"request_string" required:"false"`
}

type UpdateDivisionInput struct {
	ID       string `path:"id" minimum:"1" doc:"player ID, SteamID64"`
	Division string `path:"division"`
}
