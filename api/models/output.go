package models

type PlayerOutput struct {
	Body Player
}

type PlayersOutput struct {
	Body []Player
}

type SessionOutput struct {
	Body Session
}

type MapsOutput struct {
	Body []Map
}

type EventsOutput struct {
	Body []Event
}

type EventLeaderboardTimesOutput struct {
	Body []EventLeaderboardTime
}

type TimeWithPlayerOutput struct {
	Body TimeWithPlayer
}

type TimeWithLeaderboardOutput struct {
	Body TimeWithLeaderboard
}

type TimesWithPlayerOutput struct {
	Body []TimeWithPlayer
}

type EventWithLeaderboardsOutput struct {
	Body EventWithLeaderboards
}

type EventsWithLeaderboardsOutput struct {
	Body []EventWithLeaderboards
}

type RequestsWithPlayerOutput struct {
	Body []RequestWithPlayer
}

type TimeslotInfoOutput struct {
	Body TimeslotInfo
}
