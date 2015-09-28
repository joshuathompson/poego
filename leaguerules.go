package poego

//GetLeagueRules a list of all the possible rules for a league.
func (p *Poego) GetLeagueRules() (rules []Rule, err error) {

	r := p.buildRequest("GET", "/league-rules", nil)
	err = p.makeRequest(r, &rules)

	if err != nil {
		return rules, err
	}

	return rules, err
}

//GetLeagueRule returns a specific league rule for the supplied id.
func (p *Poego) GetLeagueRule(id string) (rule Rule, err error) {

	r := p.buildRequest("GET", "/league-rules/"+id, nil)
	err = p.makeRequest(r, &rule)

	if err != nil {
		return rule, err
	}

	return rule, err
}