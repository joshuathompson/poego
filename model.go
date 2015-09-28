package poego

import (
	"time"
)

type Ladder struct {
	Total   int16   `json:"total"`
	Entries Entries `json:"entries"`
}

type Entries struct {
	Online    bool      `json:"online"`
	Rank      int16     `json:"rank"`
	Dead      bool      `json:"dead"`
	Character Character `json:"character"`
	Account   Account   `json:"account"`
}

type Character struct {
	Name       string `json:"name"`
	Level      int8   `json:"level"`
	Class      string `json:"class"`
	Experience int    `json:"experience"`
}

type AccountChallenges struct {
	Total int16 `json:"total"`
}

type AccountTwitch struct {
	Name string `json:"name"`
}

type Account struct {
	Name       string            `json:"name"`
	Challenges AccountChallenges `json:"challenges,omitempty"`
	Twitch     AccountTwitch     `json:"twitch,omitempty"`
}

type Rule struct {
	Id          int16  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type League struct {
	Id          string     `json:"id"`
	Description string     `json:"description,omitempty"`
	RegisterAt  *time.Time `json:"registerAt,omitempty"`
	Event       bool       `json:"event,omitempty"`
	Url         string     `json:"url"`
	StartAt     *time.Time `json:"startAt"`
	EndAt       *time.Time `json:"endAt"`
	Rules       []Rule     `json:"rules,omitempty"`
	Ladder      Ladder     `json:"ladder,omitempty"`
}

type PvpMatch struct {
	Id            string     `json:"id"`
	StartAt       *time.Time `json:"startAt"`
	EndAt         *time.Time `json:"endAt"`
	Url           string     `json:"url"`
	Description   string     `json:"description"`
	GlickoRatings bool       `json:"glickoRatings"`
	Pvp           bool       `json:"pvp"`
	Style         string     `json:"style"`
	RegisterAt    *time.Time `json:"registerAt"`
}

type ApiError struct {
	Error struct {
		Code    int8   `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
