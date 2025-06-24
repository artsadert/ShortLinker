package entities

import (
	"testing"
)

func TestLinkEmpty(t *testing.T) {
	link := NewLink("google.com")
	res := link.Shorter()
	if res == "" {
		t.Error("link is empty")
	}
	t.Logf("Success: %s", res)
}
