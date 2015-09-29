package poego

import (
	"net/url"
	"testing"
)

func TestGetLeagues(t *testing.T) {

	p := NewPoeApi()

	_, err := p.GetLeagues(nil)

	if err != nil {
		t.Error("Expected Leagues JSON, got: %+v", err)
	}
}

func TestGetLeague(t *testing.T) {

	p := NewPoeApi()

	v := url.Values{}
	v.Add("ladder", "1")

	l, err := p.GetLeague("Hardcore", v)

	if err != nil {
		t.Error("Expected League JSON, got: %+v", err)
	}

	if l.Ladder == nil {
		t.Error("Expected Ladder to exist in response as ladder optional parameter was set to 1.")
	}
}
