package models

type PlayerIDInput struct {
	PlayerID string `path:"player_id" doc:"player id, SteamID64"`
}

type UpdatePlayerDivInput struct {
	PlayerID    string `path:"player_id" doc:"player id, SteamID64"`
	PlayerClass string `path:"player_class" enum:"Soldier,Demo"`
	Div         string `path:"div"`
}

type UpdatePlayerAliasInput struct {
	PlayerID string `path:"player_id" doc:"player id, SteamID64"`
	Alias    string `path:"alias"`
}

type TempusIDInput struct {
	TempusID int64 `path:"tempus_id" doc:"see: https://tempus2.xyz/"`
}

type TradeTokenInput struct {
	SteamTradeURL string `path:"steam_trade_url"`
}

type EventIDInput struct {
	ID int64 `path:"event_id"`
}

type EventKindInput struct {
	Kind string `path:"event_kind"`
}

type EventKindAndIDInput struct {
	Kind   string `path:"event_kind"`
	KindID int64  `path:"kind_id"`
}

type EventInput struct {
	Body Event
}

type LeaderboardsInput struct {
	Body []Leaderboard
}

type LeaderboardIDInput struct {
	ID int64 `path:"leaderboard_id"`
}

type UnverifiedTimeInput struct {
	ID      int64  `path:"leaderboard_id"`
	RunTime string `path:"run_time"`
}

type PlayerTimeInput struct {
	LeaderboardID int64   `path:"leaderboard_id"`
	PlayerID      string  `path:"player_id"`
	Duration      float64 `path:"duration"`
}

type TimeIDInput struct {
	ID int64 `path:"time_id"`
}

type PlayerClassInput struct {
	PlayerClass string `path:"player_class" enum:"Soldier,Demo"`
}

type MapNameInput struct {
	MapName string `path:"map_name"`
}

type LauncherInput struct {
	Launcher string `path:"launcher" enum:"stock,original,mangler,none"`
}

type RequestInput struct {
	Kind    string `path:"request_kind" enum:"alias update,soldier div,demo div"`
	Content string `path:"content"`
}

type RequestIDInput struct {
	ID int64 `path:"request_id"`
}

type TimeslotIDInput struct {
	ID int64 `path:"timeslot_id"`
}

type TimeslotInput struct {
	Body MOTWTimeslot
}
