package restContext

import (
	"github.com/golobby/container"
	"github.com/halprin/email-conceal/manager/controllers"
	"github.com/halprin/email-conceal/manager/usecases"
)

func Init() {
	container.Singleton(func() controllers.ConcealEmailController {
		return controllers.ConcealEmailController{}
	})

	container.Singleton(func() usecases.ConcealEmailUsecase {
		return usecases.ConcealEmailUsecase{}
	})
}
