package localFileWatch

import (
	"github.com/halprin/email-conceal/src/context"
	forwardEmailController "github.com/halprin/email-conceal/src/controllers/forwardEmail"
	"github.com/halprin/email-conceal/src/gateways/dynamodb"
	"github.com/halprin/email-conceal/src/gateways/localFileReader"
	"github.com/halprin/email-conceal/src/gateways/localFileWriteEmail"
	"github.com/halprin/email-conceal/src/gateways/osEnvironmentVariable"
	forwardEmailUsecase "github.com/halprin/email-conceal/src/usecases/forwardEmail"
)

func Init() {
	var applicationContext = context.ApplicationContext{}

	//controllers
	applicationContext.Bind(func() forwardEmailController.ForwardEmail {
		return forwardEmailController.CliForwardController{}
	})

	//usecases
	applicationContext.Bind(func() forwardEmailUsecase.ForwardEmailUsecase {
		return forwardEmailUsecase.ForwardEmailUsecaseImpl{}
	})

	//gateways
	dynamoDbGateway := dynamodb.DynamoDbGateway{}

	applicationContext.Bind(func() forwardEmailUsecase.ReadEmailGateway {
		return localFileReader.LocalFileReader{}
	})

	applicationContext.Bind(func() forwardEmailUsecase.SendEmailGateway {
		return localFileWriteEmail.LocalFileWriteEmailGateway{}
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
