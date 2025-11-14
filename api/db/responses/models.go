package responses

import database "github.com/spiritov/jump/api/db/queries"

type PlayerIDInput struct {
	ID int64 `path:"id" minimum:"1" doc:"player ID"`
}

type PlayerSteamID64Input struct {
	SteamID64 uint64 `path:"steam_id64" minimum:"76561197960265728" doc:"player SteamID64"`
}

// todo: response type for player
type PlayerOutput struct {
	Body database.Player
}

type Session struct {
	ID             int64  `json:"id"`
	DisplayName    string `json:"displayName"`
	SteamAvatarURL string `json:"steamAvatarUrl"`
}

type SessionOutput struct {
	Body Session
}
