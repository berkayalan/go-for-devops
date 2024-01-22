package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type lambdaParameters struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event *lambdaParameters) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}

	for index := 0; index < event.Count; index++ {
		fmt.Println("Hello %s!", event.Message)
	}

	message := fmt.Sprintf("Hello %s!", event.Message)

	return &message, nil
}
