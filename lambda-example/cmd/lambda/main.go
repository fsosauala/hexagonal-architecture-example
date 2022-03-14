package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fsosauala/lambda-example/internal/container"
)

func main() {
	lambdaHandler := container.Initialize()
	lambda.Start(lambdaHandler.LambdaHandler)
}
