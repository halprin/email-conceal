package restContext

import (
	"github.com/halprin/email-conceal/manager/context"
)

type RestApplicationContext struct{
	controllerSet RestApplicationContextControllers
	gatewaySet    RestApplicationContextGateways
	usecaseSet    RestApplicationContextUsecases
	librarySet    RestApplicationContextLibraries
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
	appContext.librarySet = RestApplicationContextLibraries{
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

func (appContext *RestApplicationContext) Libraries() context.ApplicationContextLibraries {
	return &appContext.librarySet
}
