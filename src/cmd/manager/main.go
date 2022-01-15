package main

import (
	awsLambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/halprin/email-conceal/src/external/manager/lambda"
	"github.com/halprin/email-conceal/src/external/manager/localRest"
	"os"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "local" {
		localRest.Init()
		localRest.Rest()
		return
	}

	lambda.Init()
	lambda.PrepareLambda()
	awsLambda.Start(lambda.LambdaHandler)
}
