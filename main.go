package main

import (
	"github.com/MikeB1124/registration-lambda/controllers"
	"github.com/aquasecurity/lmdrouter"
	"github.com/aws/aws-lambda-go/lambda"
)

var router *lmdrouter.Router

func init() {
	router = lmdrouter.NewRouter("")
	router.Route("POST", "/signup", controllers.AccountSignup)
	router.Route("POST", "/login", controllers.AccountLogin)
}

func main() {
	lambda.Start(router.Handler)
}
