package poego

import (
	"net/url"
)

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
