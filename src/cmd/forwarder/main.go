package main

import (
	awsLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/halprin/email-conceal/src/external/forwarder/lambda"
	"github.com/halprin/email-conceal/src/external/forwarder/localFileWatch"
	"os"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "local" {
		localFileWatch.Init()
		localFileWatch.LocalFileWatcher()
		return
	}

	lambda.Init()
	awsLambda.Start(lambda.LambdaHandler)
}
