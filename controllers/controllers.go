package controllers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/MikeB1124/registration-lambda/db"
	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func AccountSignup(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Signup request %+v\n", event)

	var user db.User
	err := json.Unmarshal([]byte(event.Body), &user)
	if err != nil {
		return createResponse(Response{Message: err.Error(), StatusCode: 400})
	}

	if user.Username == "" || user.Password == "" {
		return createResponse(Response{Message: "Username and password are required", StatusCode: 400})
	}

	exixts, err := db.UserExists(user)
	if err != nil {
		return createResponse(Response{Message: err.Error(), StatusCode: 500})
	}

	if exixts {
		return createResponse(Response{Message: "User already exists", StatusCode: 400})
	}

	timeZone, _ := time.LoadLocation("America/Los_Angeles")
	user.CreatedAt = time.Now().UTC().In(timeZone).Format("2006-01-02T15:04:05Z")
	user.LastLogin = ""

	if err := db.InsertNewUser(user); err != nil {
		return createResponse(Response{Message: err.Error(), StatusCode: 500})
	}

	return createResponse(Response{Message: "You have signed up!", StatusCode: 200})
}

func AccountLogin(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Login request %+v\n", event)

	var user db.User
	err := json.Unmarshal([]byte(event.Body), &user)
	if err != nil {
		return createResponse(Response{Message: err.Error(), StatusCode: 400})
	}

	if user.Username == "" || user.Password == "" {
		return createResponse(Response{Message: "Username and password are required", StatusCode: 400})
	}

	validLogin, err := db.ValidLogin(user)
	if err != nil {
		return createResponse(Response{Message: err.Error(), StatusCode: 500})
	}

	if !validLogin {
		return createResponse(Response{Message: "Invalid username or password", StatusCode: 401})
	}

	timeZone, _ := time.LoadLocation("America/Los_Angeles")
	lastLogin := time.Now().UTC().In(timeZone).Format("2006-01-02T15:04:05Z")

	if err := db.UpdateLastLogin(lastLogin, user.Username); err != nil {
		return createResponse(Response{Message: err.Error(), StatusCode: 500})
	}

	return createResponse(Response{Message: "You have logged in!", StatusCode: 200})
}
