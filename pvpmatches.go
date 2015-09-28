package poego

import (
	"net/url"
	"time"
)

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

//GetPvpMatches returns a list of all upcoming PvP matches.
//Optional value "seasonId" specifies which season of PvP matches is returned.
func (p *Poego) GetPvpMatches(v url.Values) (pvpMatches []PvpMatch, err error) {

	r := p.buildRequest("GET", "/pvp-matches", v)
	err = p.makeRequest(r, &pvpMatches)

	if err != nil {
		return pvpMatches, err
	}

	return pvpMatches, err
}
