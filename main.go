package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event LambdaEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", event.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
