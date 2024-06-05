package controllers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func createResponse(statusCode int, response Response) (events.APIGatewayProxyResponse, error) {
	responseBody, err := json.Marshal(response)
	if err != nil {
		responseBody, _ = json.Marshal(Response{Error: err.Error()})
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       string(responseBody),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(responseBody),
	}, nil
}
