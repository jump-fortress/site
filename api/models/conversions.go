package models

import "github.com/jump-fortress/site/db/queries"

func GetSessionResponse(p queries.Player) Session {
	return Session{
		ID:        p.ID,
		Role:      p.Role,
		Alias:     p.Alias.String,
		AvatarURL: p.AvatarUrl.String,
	}
}

func GetPlayerResponse(p queries.Player, sensitive bool) Player {
	tradeToken := p.TradeToken.String
	if sensitive {
		tradeToken = ""
	}

	return Player{
		ID:           p.ID,
		Role:         p.Role,
		Alias:        p.Alias.String,
		SoldierDiv:   p.SoldierDiv.String,
		DemoDiv:      p.DemoDiv.String,
		AvatarURL:    p.AvatarUrl.String,
		TradeToken:   tradeToken,
		TempusID:     p.TempusID.Int64,
		Country:      p.Country.String,
		CountryCode:  p.CountryCode.String,
		ClassPref:    p.ClassPref,
		MapPref:      p.MapPref.String,
		LauncherPref: p.LauncherPref.String,
		CreatedAt:    p.CreatedAt,
	}
}

func GetEventWithLeaderboardsResponse(els []queries.SelectEventLeaderboardsRow, sensitive bool) EventWithLeaderboards {
	e := els[0].Event
	ewl := EventWithLeaderboards{
		Event: Event{
			ID:          e.ID,
			Kind:        e.Kind,
			KindID:      e.KindID,
			PlayerClass: e.Class,
			VisibleAt:   e.VisibleAt,
			StartsAt:    e.StartsAt,
			EndsAt:      e.EndsAt,
			CreatedAt:   e.CreatedAt,
		},
		Leaderboards: []Leaderboard{},
	}

	for _, el := range els {
		l := el.Leaderboard
		if sensitive {
			l.Map = ""
		}
		ewl.Leaderboards = append(ewl.Leaderboards, Leaderboard{
			ID:      l.ID,
			EventID: l.EventID,
			Div:     l.Div.String,
			Map:     l.Map,
		})
	}

	return ewl
}

func GetMapResponse(m queries.Map) Map {
	return Map{
		Name:          m.Name,
		Courses:       m.Courses.Int64,
		Bonuses:       m.Bonuses.Int64,
		SoldierTier:   m.SoldierTier,
		DemoTier:      m.DemoTier,
		SoldierRating: m.SoldierRating,
		DemoRating:    m.DemoRating,
	}
}

func GetTimeResponse(t queries.Time) Time {
	return Time{
		ID:            t.ID,
		LeaderboardID: t.LeaderboardID,
		PlayerID:      t.PlayerID,
		TempusTimeID:  t.TempusTimeID.Int64,
		Duration:      t.Duration,
		Verified:      t.Verified,
		CreatedAt:     t.CreatedAt,
	}
}
