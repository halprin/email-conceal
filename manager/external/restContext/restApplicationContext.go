package restContext

import (
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/external/lib"
	"os"
)

type RestApplicationContext struct{
	controllerSet RestApplicationContextControllers
	gatewaySet    RestApplicationContextGateways
	usecaseSet    RestApplicationContextUsecases
}

func NewRestApplicationContext() *RestApplicationContext {
	appContext := &RestApplicationContext{}

	appContext.controllerSet = RestApplicationContextControllers{
		ParentContext: appContext,
	}
	appContext.gatewaySet = RestApplicationContextGateways{
		ParentContext: appContext,
	}
	appContext.usecaseSet = RestApplicationContextUsecases{
		ParentContext: appContext,
	}

	return appContext
}

func (appContext *RestApplicationContext) Controllers() context.ApplicationContextControllers {
	return &appContext.controllerSet
}

func (appContext *RestApplicationContext) Gateways() context.ApplicationContextGateways {
	return &appContext.gatewaySet
}

func (appContext *RestApplicationContext) Usecases() context.ApplicationContextUsecases {
	return &appContext.usecaseSet
}

func (appContext *RestApplicationContext) GenerateRandomUuid() string {
	return lib.GenerateGoogleRandomUuid(appContext)
}

func (appContext *RestApplicationContext) Exit(returnCode int) {
	os.Exit(returnCode)
}
