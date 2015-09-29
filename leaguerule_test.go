package poego

import (
	"testing"
)

func TestGetLeagueRules(t *testing.T) {

	p := NewPoeApi()

	_, err := p.GetLeagueRules()

	if err != nil {
		t.Error("Expected League Rules JSON, got: %+v", err)
	}
}

func TestGetLeagueRule(t *testing.T) {

	p := NewPoeApi()

	_, err := p.GetLeagueRule("4")

	if err != nil {
		t.Error("Expected League Rule JSON, got: %+v", err)
	}
}
