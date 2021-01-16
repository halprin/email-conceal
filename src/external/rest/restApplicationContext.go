package rest

import (
	"github.com/halprin/email-conceal/src/context"
	concealEmailController "github.com/halprin/email-conceal/src/controllers/concealEmail"
	"github.com/halprin/email-conceal/src/external/lib"
	"github.com/halprin/email-conceal/src/gateways/dynamodb"
	"github.com/halprin/email-conceal/src/gateways/osEnvironmentVariable"
	concealEmailUsecase "github.com/halprin/email-conceal/src/usecases/concealEmail"
	actualEmailUsecase "github.com/halprin/email-conceal/src/usecases/actualEmail"
	actualEmailController "github.com/halprin/email-conceal/src/controllers/actualEmail"
)

func init() {
	applicationContext := context.ApplicationContext{}

	//controllers
	concealEmailControllerInstance := concealEmailController.ConcealEmailController{}
	actualEmailControllerInstance := actualEmailController.ActualEmailController{}

	applicationContext.Bind(func() concealEmailController.ConcealEmailController {
		return concealEmailControllerInstance
	})

	applicationContext.Bind(func() actualEmailController.ActualEmailController {
		return actualEmailControllerInstance
	})

	//usecases
	concealEmailUsecaseInstance := concealEmailUsecase.ConcealEmailUsecaseImpl{}
	actualEmailUsecaseInstance := actualEmailUsecase.ActualEmailUsecaseImpl{}

	applicationContext.Bind(func() concealEmailUsecase.ConcealEmailUsecase {
		return concealEmailUsecaseInstance
	})

	applicationContext.Bind(func() actualEmailUsecase.ActualEmailUsecase {
		return actualEmailUsecaseInstance
	})

	//gateways
	dynamoDbGateway := dynamodb.DynamoDbGateway{}
	environmentGateway := osEnvironmentVariable.OsEnvironmentGateway{}

	applicationContext.Bind(func() concealEmailUsecase.ConcealEmailGateway {
		return dynamoDbGateway
	})

	applicationContext.Bind(func() context.EnvironmentGateway {
		return environmentGateway
	})

	//libraries
	googleUuid := lib.GoogleUuid{}

	applicationContext.Bind(func() context.UuidLibrary {
		return googleUuid
	})

	//inits
	actualEmailControllerInstance.Init()
	dynamoDbGateway.Init()
}
