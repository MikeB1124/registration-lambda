package controllers

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func AccountSignup(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %+v\n", event)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "You have signed up!",
	}, nil
}

func AccountLogin(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %+v\n", event)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "You have logged in!",
	}, nil
}
