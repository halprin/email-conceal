package testApplicationContext

import "github.com/halprin/email-conceal/manager/context"

type TestApplicationContext struct {
	controllerSet TestApplicationContextControllers

	ReceivedEnvironmentGatewayArguments string
	ReturnFromEnvironmentGateway        map[string]string

	ReceivedAddConcealedEmailToActualEmailMappingGatewayConcealPrefixArgument string
	ReceivedAddConcealedEmailToActualEmailMappingGatewayEmailArgument         string
	ReturnErrorFromAddConcealedEmailToActualEmailMappingGateway               error

	ReceivedDeleteConcealedEmailToActualEmailMappingGatewayConcealPrefixArgument string
	ReturnErrorFromDeleteConcealedEmailToActualEmailMappingGateway               error

	ReceivedConcealEmailUsecaseEmail   string
	ReturnFromConcealEmailUsecase      string
	ReturnErrorFromConcealEmailUsecase error

	ReceivedDeleteConcealEmailUsecaseConcealPrefixArgument string
	ReturnErrorFromDeleteConcealEmailUsecase               error

	ReturnFromGenerateRandomUuid string

	ReceivedExitReturnCode int
}

func (appContext *TestApplicationContext) Controllers() context.ApplicationContextControllers {
	return &appContext.controllerSet
}

func (appContext *TestApplicationContext) EnvironmentGateway(key string) string {
	appContext.ReceivedEnvironmentGatewayArguments = key
	return appContext.ReturnFromEnvironmentGateway[key]
}

func (appContext *TestApplicationContext) AddConcealedEmailToActualEmailMappingGateway(concealPrefix string, actualEmail string) error {
	appContext.ReceivedAddConcealedEmailToActualEmailMappingGatewayConcealPrefixArgument = concealPrefix
	appContext.ReceivedAddConcealedEmailToActualEmailMappingGatewayEmailArgument = actualEmail
	return appContext.ReturnErrorFromAddConcealedEmailToActualEmailMappingGateway
}

func (appContext *TestApplicationContext) DeleteConcealedEmailToActualEmailMappingGateway(concealPrefix string) error {
	appContext.ReceivedDeleteConcealedEmailToActualEmailMappingGatewayConcealPrefixArgument = concealPrefix
	return appContext.ReturnErrorFromDeleteConcealedEmailToActualEmailMappingGateway
}

func (appContext *TestApplicationContext) AddConcealEmailUsecase(email string) (string, error) {
	appContext.ReceivedConcealEmailUsecaseEmail = email
	return appContext.ReturnFromConcealEmailUsecase, appContext.ReturnErrorFromConcealEmailUsecase
}

func (appContext *TestApplicationContext) DeleteConcealEmailUsecase(concealPrefix string) error {
	appContext.ReceivedDeleteConcealEmailUsecaseConcealPrefixArgument = concealPrefix
	return appContext.ReturnErrorFromDeleteConcealEmailUsecase
}

func (appContext *TestApplicationContext) GenerateRandomUuid() string {
	return appContext.ReturnFromGenerateRandomUuid
}

func (appContext *TestApplicationContext) Exit(returnCode int) {
	appContext.ReceivedExitReturnCode = returnCode
}
