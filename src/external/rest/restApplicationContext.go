package rest

import (
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/controllers/concealEmail"
	"github.com/halprin/email-conceal/src/external/lib"
	"github.com/halprin/email-conceal/src/gateways/dynamodb"
	"github.com/halprin/email-conceal/src/gateways/osEnvironmentVariable"
	"github.com/halprin/email-conceal/src/usecases"
)

func init() {
	var applicationContext = context.ApplicationContext{}

	//controllers
	applicationContext.Bind(func() concealEmail.ConcealEmailController {
		return concealEmail.ConcealEmailController{}
	})

	//usecases
	applicationContext.Bind(func() usecases.ConcealEmailUsecase {
		return usecases.ConcealEmailUsecaseImpl{}
	})

	//gateways
	applicationContext.Bind(func() usecases.ConcealEmailGateway {
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
