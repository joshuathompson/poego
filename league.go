package poego

import (
	"net/url"
	"time"
)

type League struct {
	Id          string     `gorethink:"id" json:"id"`
	Description string     `gorethink:"description,omitempty" json:"description,omitempty"`
	RegisterAt  *time.Time `gorethink:"registerAt,omitempty" json:"registerAt,omitempty"`
	Event       bool       `gorethink:"event,omitempty" json:"event,omitempty"`
	Url         string     `gorethink:"url" json:"url"`
	StartAt     *time.Time `gorethink:"startAt" json:"startAt"`
	EndAt       *time.Time `gorethink:"endAt" json:"endAt"`
	Rules       []Rules    `gorethink:"rules,omitempty" json:"rules,omitempty"`
	Ladder      *Ladder    `gorethink:"ladder,omitempty" json:"ladder,omitempty"`
}

//GetLeagues returns a list of current leagues.
//Optional value "type" can be set to all, event, or seasonal in orer to filter results to the specified type.
//Optional value "season" is required when "type" is set to season and specifies the id of the season to get.
//Optional value "compact" can be set to 0 or 1 specifies whether to include rules/minimal ladder values
//and effects the total number of results returned as it increases the total volume of JSON.
func (p *Poego) GetLeagues(v url.Values) (leagues []League, err error) {

	r := p.buildRequest("GET", "/leagues", v)
	err = p.makeRequest(r, &leagues)

	if err != nil {
		return leagues, err
	}

	return leagues, err
}

//GetLeague returns a league for the supplied id.
//Optional value "ladder" can be set to 0 (default) or 1 and specifies whether to include ladder data in the response.
//Optional value "ladderLimit" should be a value between 20 and 200 and specifies the number of ladder results to return.
//Optional value "ladderOffset" should either be 0 or 1 and specifies where to begin retrieving ladder values.
//Optional values "ladderLimit" and "ladderOffset" should only be used when "ladder" is set to 1.
func (p *Poego) GetLeague(id string, v url.Values) (league League, err error) {

	r := p.buildRequest("GET", "/leagues/"+id, v)
	err = p.makeRequest(r, &league)

	if err != nil {
		return league, err
	}

	return league, err
}

//GetEntireLeagueLadder returns a league for the supplied id and includes all 15000 users in the ladder.
//This will take roughly 25 seconds to run as explained in the GetEntireLadder function.
func (p *Poego) GetEntireLeagueLadder(id string) (league League, err error) {

	league, err = p.GetLeague(id, nil)

	if err != nil {
		return league, err
	}

	ladder, err := p.GetEntireLadder(id)

	if err != nil {
		return league, err
	}

	league.Ladder = &ladder

	return league, err
}
