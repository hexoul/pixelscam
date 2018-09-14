package core

import (
	"testing"

	"github.com/hexoul/pixelscam/json"
)

func TestRoute(t *testing.T) {
	var req json.Request
	req.ID = 100
	req.Method = HeartBeat
	if ret := Route(req); !ret {
		t.Fatalf("Failed to route")
	}
}
