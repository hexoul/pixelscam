package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	respBody := ""
	statusCode := 200

	return events.APIGatewayProxyResponse{Body: respBody, StatusCode: statusCode}, nil
}

func main() {
	lambda.Start(lambdaHandler)
}
