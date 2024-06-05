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
	Error   string `json:"error"`
	Message string `json:"message"`
}

func AccountSignup(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Signup request %+v\n", event)

	var user db.User
	err := json.Unmarshal([]byte(event.Body), &user)
	if err != nil {
		return createResponse(400, Response{Error: err.Error()})
	}

	if user.Username == "" || user.Password == "" {
		return createResponse(400, Response{Error: "Username and password are required"})
	}

	exixts, err := db.UserExists(user)
	if err != nil {
		return createResponse(500, Response{Error: err.Error()})
	}

	if exixts {
		return createResponse(400, Response{Error: "User already exists"})
	}

	timeZone, _ := time.LoadLocation("America/Los_Angeles")
	user.CreatedAt = time.Now().UTC().In(timeZone).Format("2006-01-02T15:04:05Z")
	user.LastLogin = ""

	if err := db.InsertNewUser(user); err != nil {
		return createResponse(500, Response{Error: err.Error()})
	}

	return createResponse(200, Response{Message: "You have signed up!"})
}

func AccountLogin(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Login request %+v\n", event)

	var user db.User
	err := json.Unmarshal([]byte(event.Body), &user)
	if err != nil {
		return createResponse(400, Response{Error: err.Error()})
	}

	if user.Username == "" || user.Password == "" {
		return createResponse(400, Response{Error: "Username and password are required"})
	}

	validLogin, err := db.ValidLogin(user)
	if err != nil {
		return createResponse(500, Response{Error: err.Error()})
	}

	if !validLogin {
		return createResponse(400, Response{Error: "Invalid username or password"})
	}

	timeZone, _ := time.LoadLocation("America/Los_Angeles")
	lastLogin := time.Now().UTC().In(timeZone).Format("2006-01-02T15:04:05Z")

	if err := db.UpdateLastLogin(lastLogin, user.Username); err != nil {
		return createResponse(500, Response{Error: err.Error()})
	}

	return createResponse(200, Response{Message: "You have logged in!"})
}
