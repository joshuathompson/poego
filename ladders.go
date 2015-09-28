package poego

import (
	"net/url"
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

type Account struct {
	Name       string            `json:"name"`
	Challenges AccountChallenges `json:"challenges,omitempty"`
	Twitch     AccountTwitch     `json:"twitch,omitempty"`
}

type AccountChallenges struct {
	Total int16 `json:"total"`
}

type AccountTwitch struct {
	Name string `json:"name"`
}

//GetLadder returns a ladder for the supplied id.
//Optional value "limit" which is an integer between 20 and 200 specifies how many entries to pull.
//Optional value "offset" specifies how far from 0 to begin pulling entries.
//The max amount of entries in a ladder is 15000 which means you need to make 75 requests to pull an entire set.
func (p *Poego) GetLadder(id string, v url.Values) (ladder []Ladder, err error) {

	r := p.buildRequest("GET", "/ladders/"+id, v)
	err = p.makeRequest(r, &ladder)

	if err != nil {
		return ladder, err
	}

	return ladder, err
}
