package models

type TempusPlayer struct {
	Name        string `json:"name"`
	ID          int64  `json:"id"`
	SteamID     string `json:"steamid"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

// endpoint: /players/id/{playerId}/stats
type TempusPlayerInfo struct {
	PlayerInfo TempusPlayer `json:"player_info"`
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

type TempusMap struct {
	ID     int64        `json:"id"`
	Name   string       `json:"name"`
	Zones  TempusZones  `json:"zone_counts"`
	Tier   TempusTier   `json:"tier_info"`
	Rating TempusRating `json:"rating_info"`
}

type TempusMaps struct {
	Maps []TempusMap
}

type TempusTime struct {
	ID       int64   `json:"id"`
	Date     float64 `json:"date"`
	Duration float64 `json:"duration"`
}

type TempusTimeInfo struct {
	PR TempusTime `json:"result"`
}
