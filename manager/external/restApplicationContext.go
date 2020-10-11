package external

import (
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/controllers"
	"github.com/halprin/email-conceal/manager/external/lib"
	"github.com/halprin/email-conceal/manager/gateways"
	"github.com/halprin/email-conceal/manager/usecases"
)

func init() {
	var applicationContext = context.ApplicationContext{}

	//controllers
	applicationContext.Bind(func() controllers.ConcealEmailController {
		return controllers.ConcealEmailController{}
	})

	//usecases
	applicationContext.Bind(func() usecases.ConcealEmailUsecase {
		return usecases.ConcealEmailUsecaseImpl{}
	})

	//gateways
	applicationContext.Bind(func() usecases.ConcealEmailGateway {
		return gateways.DynamoDbGateway{}
	})

	applicationContext.Bind(func() context.EnvironmentGateway {
		return gateways.OsEnvironmentGateway{}
	})

	//libraries
	applicationContext.Bind(func() context.UuidLibrary {
		return lib.GoogleUuid{}
	})
}
