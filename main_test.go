package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestEntrance(t *testing.T) {
	event := events.APIGatewayProxyRequest{}
	_, err := lambdaHandler(nil, event)
	if err != nil {
		t.Errorf("Failed to enter")
	}
}
