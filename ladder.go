package poego

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Ladder struct {
	Total   int       `gorethink:"total" json:"total"`
	Entries []Entries `gorethink:"entries" json:"entries"`
}

type Entries struct {
	Online    bool       `gorethink:"online" json:"online"`
	Rank      int        `gorethink:"rank" json:"rank"`
	Dead      bool       `gorethink:"dead" json:"dead"`
	Character *Character `gorethink:"character" json:"character"`
	Account   *Account   `gorethink:"account" json:"account"`
}

type Character struct {
	Name       string `gorethink:"name" json:"name"`
	Level      int    `gorethink:"level" json:"level"`
	Class      string `gorethink:"class" json:"class"`
	Experience int    `gorethink:"experience" json:"experience"`
}

type Account struct {
	Name       string             `gorethink:"name" json:"name"`
	Challenges *AccountChallenges `gorethink:"challenges,omitempty" json:"challenges,omitempty"`
	Twitch     *AccountTwitch     `gorethink:"twitch,omitempty" json:"twitch,omitempty"`
}

type AccountChallenges struct {
	Total int `gorethink:"total" json:"total"`
}

type AccountTwitch struct {
	Name string `gorethink:"name" json:"name"`
}

//GetLadder returns a ladder for the supplied id.
//Optional value "limit" which is an integer between 20 and 200 specifies how many entries to pull.
//Optional value "offset" specifies how far from 0 to begin pulling entries.
//The max amount of entries in a ladder is 15000 which means you need to make 75 requests to pull an entire set.
func (p *Poego) GetLadder(id string, v url.Values) (ladder Ladder, err error) {

	r := p.buildRequest("GET", "/ladders/"+id, v)
	err = p.makeRequest(r, &ladder)

	if err != nil {
		return ladder, err
	}

	return ladder, err
}

func (p *Poego) GetEntireLadder(id string) (l Ladder, e error) {

	requests := p.buildRequestsForEntireLadder("GET", "/ladders/"+id)
	ladders := make(chan Ladder, len(requests))
	errs := make(chan error, len(requests))

	throttle := time.Tick(time.Second)
	for _, req := range requests {
		<-throttle
		go func(req *http.Request) {

			var ladder Ladder
			err := p.makeRequest(req, &ladder)

			if err != nil {
				errs <- err
			}

			ladders <- ladder
		}(req)
	}

	for i := 0; i < len(requests); i++ {
		select {
		case ladder := <-ladders:
			l.Total = ladder.Total
			l.Entries = append(l.Entries, ladder.Entries...)
		case err := <-errs:
			return l, err
		}
	}

	return l, e
}
