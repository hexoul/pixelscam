package json

import (
	"testing"
)

func TestSerialize(t *testing.T) {
	testMsg := "{\"method\":\"func1\",\"params\":[\"a\",1],\"id\":100}"
	if reqRet := GetRequestFromJSON(testMsg); reqRet.ID != 100 {
		t.Errorf("Failed to serialize Request")
	}
}
