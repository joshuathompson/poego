package poego

import (
	"net/url"
	"time"
)

type League struct {
	Id          string     `json:"id"`
	Description string     `json:"description,omitempty"`
	RegisterAt  *time.Time `json:"registerAt,omitempty"`
	Event       bool       `json:"event,omitempty"`
	Url         string     `json:"url"`
	StartAt     *time.Time `json:"startAt"`
	EndAt       *time.Time `json:"endAt"`
	Rules       []Rules    `json:"rules,omitempty"`
	Ladder      *Ladder    `json:"ladder,omitempty"`
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
