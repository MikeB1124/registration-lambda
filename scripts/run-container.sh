#!/bin/bash

docker stop registration-lambda
docker rm registration-lambda
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go
docker build -t registration-lambda-image .
docker run --name registration-lambda -p 9000:8080 --env-file .env registration-lambda-image