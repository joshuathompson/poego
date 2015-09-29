package poego

import (
	"net/url"
	"testing"
)

func TestGetLadder(t *testing.T) {

	p := NewPoeApi()

	v := url.Values{}
	v.Add("offset", "1")
	v.Add("limit", "2")

	l, err := p.GetLadder("Standard", v)

	if err != nil {
		t.Error("Expected Ladder JSON, got: %+v", err)
	}

	if len(l.Entries) > 2 {
		t.Error("Expected Ladder length of %d and got %d", 2, len(l.Entries))
	}
}

func TestGetLadderInvalidQuery(t *testing.T) {

	p := NewPoeApi()

	v := url.Values{}
	v.Add("offset", "1")
	v.Add("limit", "2000")

	_, err := p.GetLadder("Standard", v)

	if err == nil {
		t.Error("Expected Invalid query error because limit %d is geater than max 200.", 2000)
	}
}
