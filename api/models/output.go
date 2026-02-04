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

type EventLeaderboardTimeOutput struct {
	Body EventLeaderboardTime
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
