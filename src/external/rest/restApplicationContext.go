package rest

import (
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/controllers/concealEmail"
	"github.com/halprin/email-conceal/src/external/lib"
	"github.com/halprin/email-conceal/src/gateways/dynamodb"
	"github.com/halprin/email-conceal/src/gateways/osEnvironmentVariable"
	concealEmail2 "github.com/halprin/email-conceal/src/usecases/concealEmail"
)

func init() {
	var applicationContext = context.ApplicationContext{}

	//controllers
	applicationContext.Bind(func() concealEmail.ConcealEmailController {
		return concealEmail.ConcealEmailController{}
	})

	//usecases
	applicationContext.Bind(func() concealEmail2.ConcealEmailUsecase {
		return concealEmail2.ConcealEmailUsecaseImpl{}
	})

	//gateways
	applicationContext.Bind(func() concealEmail2.ConcealEmailGateway {
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
