package poego

import (
	"net/http"
	"net/url"
	"strconv"
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
func (p *Poego) GetLadder(id string, v url.Values) (ladder Ladder, err error) {

	r := p.buildRequest("GET", "/ladders/"+id, v)
	err = p.makeRequest(r, &ladder)

	if err != nil {
		return ladder, err
	}

	return ladder, err
}

//GetEntireLadder gets up to 15000 ladder entries (api limitation) for the supplied id by making 75 requests
//at a rate of 5 requests per second.  This means that it will take roughly 15 seconds in total to execute.
//Exercise caution due to rate limits.
func (p *Poego) GetEntireLadder(id string) (l Ladder, e error) {

	var requests []*http.Request
	numUrls := TotalPossibleLadderEntries / MaxLadderSegmentSize
	ladders := make(chan Ladder, len(requests))
	errs := make(chan error, len(requests))

	for i := 0; i < numUrls; i++ {

		v := url.Values{}
		v.Add("limit", "200")
		v.Add("offset", strconv.Itoa(i*200))

		requests = append(requests, p.buildRequest("GET", "/ladders/"+id, v))
	}

	//3 requests per second
	rate := time.Tick(time.Second / 5)
	for _, req := range requests {
		go func(req *http.Request) {

			var ladder Ladder
			err := p.makeRequest(req, &ladder)

			if err != nil {
				errs <- err
			}

			ladders <- ladder
		}(req)
		<-rate
	}

	for i := 0; i < len(requests); i++ {
		select {
		case ladder := <-ladders:
			l.Total = ladder.Total
			l.Entries = append(l.Entries, ladder.Entries...)
		case err := <-errs:
			e = err
			return l, e
		}
	}

	return l, e
}
