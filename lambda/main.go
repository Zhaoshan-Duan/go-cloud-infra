package main

// Backend: handle all the backend logic for the application
// 	pakcaged and zipped to be deployed to AWS Lambda function

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)



type MyEvent struct {
	Username string `json:"username"`
}

// Take in a payload, process it, and return a response
func HandlerRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username is cannot be empty")
	}
	
	return fmt.Sprintf("Succesfully called by - %s", event.Username), nil
}


func main() {
	// On invocation, call the HandlerRequest function
	lambda.Start(HandlerRequest)
}
