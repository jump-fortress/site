package responses

import (
	"time"
)

type PlayerIDInput struct {
	ID int64 `path:"id" minimum:"1" doc:"player ID"`
}

type PlayerSteamID64Input struct {
	SteamID64 uint64 `path:"steam_id64" minimum:"76561197960265728" doc:"player SteamID64"`
}

type Player struct {
	ID              int64     `json:"id"`
	Role            string    `json:"role"`
	SteamAvatarUrl  string    `json:"steam_avatar_url"`
	TempusID        int64     `json:"tempus_id,omitempty"`
	DisplayName     string    `json:"display_name"`
	SoldierDivision string    `json:"soldier_division,omitempty"`
	DemoDivision    string    `json:"demo_division,omitempty"`
	PreferredClass  string    `json:"preferred_class"`
	CreatedAt       time.Time `json:"created_at"`
}

type PlayerOutput struct {
	Body Player
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
	ID             int64  `json:"id"`
	Role           string `json:"role"`
	DisplayName    string `json:"displayName"`
	SteamAvatarURL string `json:"steamAvatarUrl"`
}

type SessionOutput struct {
	Body Session
}
