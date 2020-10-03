package restContext

import (
	"github.com/golobby/container"
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/controllers"
	"github.com/halprin/email-conceal/manager/external/lib"
	"github.com/halprin/email-conceal/manager/gateways"
	"github.com/halprin/email-conceal/manager/usecases"
)

func Init() {
	//controllers
	container.Singleton(func() controllers.ConcealEmailController {
		return controllers.ConcealEmailController{}
	})

	//usecases
	container.Singleton(func() usecases.ConcealEmailUsecase {
		return usecases.ConcealEmailUsecase{}
	})

	//gateways
	container.Singleton(func() usecases.ConcealEmailGateway {
		return gateways.DynamoDbGateway{}
	})

	container.Singleton(func() context.EnvironmentGateway {
		return gateways.OsEnvironmentGateway{}
	})

	//libraries
	container.Singleton(func() context.UuidLibrary {
		return lib.GoogleUuid{}
	})
}
