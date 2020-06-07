package restContext

import (
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/controllers"
)

type RestApplicationContextControllers struct{
	ParentContext context.ApplicationContext
}

func (appContextControllers *RestApplicationContextControllers) ConcealEmail(arguments map[string]interface{}) (int, map[string]string) {
	return controllers.HttpConcealEmailController(arguments, appContextControllers.ParentContext)
}

func (appContextControllers *RestApplicationContextControllers) DeleteConcealEmail(arguments map[string]interface{}) (int, map[string]string) {
	return controllers.HttpDeleteConcealEmailController(arguments, appContextControllers.ParentContext)
}
