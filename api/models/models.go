package models

import "time"

var Divs = []string{"Diamond", "Platinum", "Gold", "Silver", "Bronze", "Steel", "Wood"}
var EventKinds = []string{"monthly", "archive"}

type Session struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	Alias     string `json:"alias"`
	AvatarURL string `json:"avatar_url"`
}

type Player struct {
	ID           string    `json:"id"`
	Role         string    `json:"role"`
	Alias        string    `json:"alias"`
	SoldierDiv   string    `json:"soldier_div,omitempty"`
	DemoDiv      string    `json:"demo_div,omitempty"`
	AvatarURL    string    `json:"avatar_url"`
	TradeToken   string    `json:"trade_token,omitempty"`
	TempusID     int64     `json:"tempus_id,omitempty"`
	Country      string    `json:"country,omitempty"`
	CountryCode  string    `json:"country_code,omitempty"`
	ClassPref    string    `json:"class_pref"`
	MapPref      string    `json:"map_pref,omitempty"`
	LauncherPref string    `json:"launcher_pref,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

type Event struct {
	ID          int64     `json:"id"`
	Kind        string    `json:"kind" enum:"monthly,archive"`
	KindID      int64     `json:"kind_id"`
	PlayerClass string    `json:"player_class" enum:"Soldier,Demo"`
	VisibleAt   time.Time `json:"visible_at"`
	StartsAt    time.Time `json:"starts_at"`
	EndsAt      time.Time `json:"ends_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type Leaderboard struct {
	ID      int64  `json:"id"`
	EventID int64  `json:"event_id"`
	Div     string `json:"div,omitempty"`
	Map     string `json:"map"`
}

type Map struct {
	Name          string `json:"name"`
	Courses       int64  `json:"courses"`
	Bonuses       int64  `json:"bonuses"`
	SoldierTier   int64  `json:"soldier_tier"`
	DemoTier      int64  `json:"demo_tier"`
	SoldierRating int64  `json:"soldier_rating"`
	DemoRating    int64  `json:"demo_rating"`
}

type Time struct {
	ID            int64     `json:"id"`
	LeaderboardID int64     `json:"leaderboard_id"`
	PlayerID      string    `json:"player_id"`
	TempusTimeID  int64     `json:"tempus_time_id,omitempty"`
	Duration      float64   `json:"duration"`
	Verified      bool      `json:"verified"`
	CreatedAt     time.Time `json:"created_at"`
}

type EventWithLeaderboards struct {
	Event        Event         `json:"event"`
	Leaderboards []Leaderboard `json:"leaderboards"`
}
