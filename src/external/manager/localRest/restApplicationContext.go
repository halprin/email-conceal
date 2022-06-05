package localRest

import (
	"github.com/halprin/email-conceal/src/context"
	accountController "github.com/halprin/email-conceal/src/controllers/account"
	actualEmailController "github.com/halprin/email-conceal/src/controllers/actualEmail"
	concealEmailController "github.com/halprin/email-conceal/src/controllers/concealEmail"
	"github.com/halprin/email-conceal/src/external/lib"
	"github.com/halprin/email-conceal/src/gateways/dynamodb"
	"github.com/halprin/email-conceal/src/gateways/localFileWriteEmail"
	"github.com/halprin/email-conceal/src/gateways/osEnvironmentVariable"
	accountUsecase "github.com/halprin/email-conceal/src/usecases/account"
	actualEmailUsecase "github.com/halprin/email-conceal/src/usecases/actualEmail"
	concealEmailUsecase "github.com/halprin/email-conceal/src/usecases/concealEmail"
	forwardEmailUsecase "github.com/halprin/email-conceal/src/usecases/forwardEmail"
)

func Init() {
	applicationContext := context.ApplicationContext{}

	//controllers
	concealEmailControllerInstance := concealEmailController.ConcealEmailController{}
	actualEmailControllerInstance := actualEmailController.ActualEmailController{}
	accountControllerInstance := accountController.AccountController{}

	applicationContext.Bind(func() concealEmailController.ConcealEmailController {
		return concealEmailControllerInstance
	})

	applicationContext.Bind(func() actualEmailController.ActualEmailController {
		return actualEmailControllerInstance
	})

	applicationContext.Bind(func() accountController.AccountController {
		return accountControllerInstance
	})

	//usecases
	concealEmailUsecaseInstance := concealEmailUsecase.ConcealEmailUsecaseImpl{}
	actualEmailUsecaseInstance := actualEmailUsecase.ActualEmailUsecaseImpl{}
	accountUsecaseInstance := accountUsecase.AccountUsecaseImpl{}

	applicationContext.Bind(func() concealEmailUsecase.ConcealEmailUsecase {
		return concealEmailUsecaseInstance
	})

	applicationContext.Bind(func() actualEmailUsecase.ActualEmailUsecase {
		return actualEmailUsecaseInstance
	})

	applicationContext.Bind(func() accountUsecase.AccountUsecase {
		return accountUsecaseInstance
	})

	//gateways
	dynamoDbGateway := dynamodb.DynamoDbGateway{}
	environmentGateway := osEnvironmentVariable.OsEnvironmentGateway{}

	applicationContext.Bind(func() concealEmailUsecase.ConcealEmailGateway {
		return dynamoDbGateway
	})

	applicationContext.Bind(func() actualEmailUsecase.ActualEmailConfigurationGateway {
		return dynamoDbGateway
	})

	applicationContext.Bind(func() forwardEmailUsecase.SendEmailGateway {
		return localFileWriteEmail.LocalFileWriteEmailGateway{}
	})

	applicationContext.Bind(func() context.EnvironmentGateway {
		return environmentGateway
	})

	applicationContext.Bind(func() accountUsecase.AccountConfigurationGateway {
		return dynamoDbGateway
	})

	//libraries
	googleUuid := lib.GoogleUuid{}

	applicationContext.Bind(func() context.UuidLibrary {
		return googleUuid
	})

	//inits
	actualEmailControllerInstance.Init()
	dynamoDbGateway.Init()
	actualEmailUsecaseInstance.Init()
	concealEmailUsecaseInstance.Init()
	concealEmailControllerInstance.Init()
	accountUsecaseInstance.Init()
	accountControllerInstance.Init()
}
