// Package core implements mining from heart beat
package core

import (
	"fmt"

	"github.com/hexoul/pixelscam/json"
)

const (
	// HeartBeat to mine
	HeartBeat = "HeartBeat"
)

// Route a given request to appropriate method
func Route(req json.Request) bool {
	switch req.Method {
	case HeartBeat:
		fmt.Printf("HeartBeat IN")
		return true
	}
	return false
}
