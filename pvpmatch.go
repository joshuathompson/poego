package poego

import (
	"net/url"
	"time"
)

type PvpMatch struct {
	Id            string     `gorethink:"id" json:"id"`
	StartAt       *time.Time `gorethink:"startAt" json:"startAt"`
	EndAt         *time.Time `gorethink:"endAt" json:"endAt"`
	Url           string     `gorethink:"url" json:"url"`
	Description   string     `gorethink:"description" json:"description"`
	GlickoRatings bool       `gorethink:"glickoRatings" json:"glickoRatings"`
	Pvp           bool       `gorethink:"pvp" json:"pvp"`
	Style         string     `gorethink:"style" json:"style"`
	RegisterAt    *time.Time `gorethink:"registerAt" json:"registerAt"`
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
