// Package json is a JSON parser for request
package json

import (
	"encoding/json"
)

// Request interface following JSON
type Request struct {
	ID     int32         `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

// GetRequestFromJSON returns Request struct from JSON string
func GetRequestFromJSON(msg string) Request {
	var data Request
	json.Unmarshal([]byte(msg), &data)
	return data
}

func (r *Request) String() string {
	if ret, err := json.Marshal(r); err == nil {
		return string(ret)
	}
	return ""
}
