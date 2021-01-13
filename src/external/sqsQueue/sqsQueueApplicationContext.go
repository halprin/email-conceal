package sqsQueue

import (
	"github.com/halprin/email-conceal/src/context"
	forwardEmailController "github.com/halprin/email-conceal/src/controllers/forwardEmail"
	"github.com/halprin/email-conceal/src/gateways/awsSesSendEmail"
	"github.com/halprin/email-conceal/src/gateways/dynamodb"
	"github.com/halprin/email-conceal/src/gateways/osEnvironmentVariable"
	"github.com/halprin/email-conceal/src/gateways/s3FileReader"
	forwardEmailUsecase "github.com/halprin/email-conceal/src/usecases/forwardEmail"
)


func init() {
	var applicationContext = context.ApplicationContext{}

	//controllers
	applicationContext.Bind(func() forwardEmailController.ForwardEmail {
		return forwardEmailController.SqsQueueForwardController{}
	})

	//usecases
	applicationContext.Bind(func() forwardEmailUsecase.ForwardEmailUsecase {
		return forwardEmailUsecase.ForwardEmailUsecaseImpl{}
	})

	//gateways
	dynamoDbGateway := dynamodb.DynamoDbGateway{}

	applicationContext.Bind(func() forwardEmailUsecase.ReadEmailGateway {
		return s3FileReader.S3FileReader{}
	})

	applicationContext.Bind(func() forwardEmailUsecase.SendEmailGateway {
		return awsSesSendEmail.AwsSesSendEmailGateway{}
	})

	applicationContext.Bind(func() forwardEmailUsecase.ConfigurationGateway {
		return dynamoDbGateway
	})

	applicationContext.Bind(func() context.EnvironmentGateway {
		return osEnvironmentVariable.OsEnvironmentGateway{}
	})

	//init
	dynamoDbGateway.Init()
}
