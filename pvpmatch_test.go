package poego

import (
	"testing"
)

func TestGetPvpMatches(t *testing.T) {

	p := NewPoeApi()

	_, err := p.GetPvpMatches(nil)

	if err != nil {
		t.Error("Expected PvP matches JSON, got: %+v", err)
	}
}
