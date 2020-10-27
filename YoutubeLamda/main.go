package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == "GET" {
		stringResponse := "GET request"
		ApiResponse := events.APIGatewayProxyResponse{Body: stringResponse, StatusCode: 200}
		return ApiResponse, nil
	} else {
		err := errors.New("Method not allowed")
		ApiResponse := events.APIGatewayProxyResponse{Body: "Method not ok", StatusCode: 502}
		return ApiResponse, err
	}
}
