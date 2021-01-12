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
	var applicationContext = context.ApplicationContext{}

	//controllers
	applicationContext.Bind(func() concealEmailController.ConcealEmailController {
		return concealEmailController.ConcealEmailController{}
	})

	applicationContext.Bind(func() actualEmailController.ActualEmailController {
		return actualEmailController.ActualEmailController{}
	})

	//usecases
	applicationContext.Bind(func() concealEmailUsecase.ConcealEmailUsecase {
		return concealEmailUsecase.ConcealEmailUsecaseImpl{}
	})

	applicationContext.Bind(func() actualEmailUsecase.ActualEmailUsecase {
		return actualEmailUsecase.ActualEmailUsecase{}
	})

	//gateways
	applicationContext.Bind(func() concealEmailUsecase.ConcealEmailGateway {
		return dynamodb.DynamoDbGateway{}
	})

	applicationContext.Bind(func() context.EnvironmentGateway {
		return osEnvironmentVariable.OsEnvironmentGateway{}
	})

	//libraries
	applicationContext.Bind(func() context.UuidLibrary {
		return lib.GoogleUuid{}
	})
}
