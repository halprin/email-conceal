package restContext

import (
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/external/lib"
	"github.com/halprin/email-conceal/manager/gateways"
	"github.com/halprin/email-conceal/manager/usecases"
	"os"
)

type RestApplicationContext struct{
	controllerSet RestApplicationContextControllers
	gatewaySet    RestApplicationContextGateways
}

func NewRestApplicationContext() *RestApplicationContext {
	appContext := &RestApplicationContext{}

	appContext.controllerSet = RestApplicationContextControllers{
		ParentContext: appContext,
	}
	appContext.gatewaySet = RestApplicationContextGateways{
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

func (appContext *RestApplicationContext) EnvironmentGateway(key string) string {
	return gateways.OsEnvEnvironmentGateway(key, appContext)
}

func (appContext *RestApplicationContext) AddConcealedEmailToActualEmailMappingGateway(concealPrefix string, actualEmail string) error {
	return gateways.AddConcealedEmailToActualEmailMapping(concealPrefix, actualEmail, appContext)
}

func (appContext *RestApplicationContext) DeleteConcealedEmailToActualEmailMappingGateway(concealPrefix string) error {
	return gateways.DeleteConcealedEmailToActualEmailMapping(concealPrefix, appContext)
}

func (appContext *RestApplicationContext) AddConcealEmailUsecase(email string) (string, error) {
	return usecases.AddConcealEmailUsecase(email, appContext)
}

func (appContext *RestApplicationContext) DeleteConcealEmailUsecase(concealPrefix string) error {
	return usecases.DeleteConcealEmailMappingUsecase(concealPrefix, appContext)
}

func (appContext *RestApplicationContext) GenerateRandomUuid() string {
	return lib.GenerateGoogleRandomUuid(appContext)
}

func (appContext *RestApplicationContext) Exit(returnCode int) {
	os.Exit(returnCode)
}
