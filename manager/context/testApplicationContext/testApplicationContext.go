package testApplicationContext

import "github.com/halprin/email-conceal/manager/context"

type TestApplicationContext struct {
	ControllerSet TestApplicationContextControllers
	GatewaySet    TestApplicationContextGateways

	ReceivedConcealEmailUsecaseEmail   string
	ReturnFromConcealEmailUsecase      string
	ReturnErrorFromConcealEmailUsecase error

	ReceivedDeleteConcealEmailUsecaseConcealPrefixArgument string
	ReturnErrorFromDeleteConcealEmailUsecase               error

	ReturnFromGenerateRandomUuid string

	ReceivedExitReturnCode int
}

func (appContext *TestApplicationContext) Controllers() context.ApplicationContextControllers {
	return &appContext.ControllerSet
}

func (appContext *TestApplicationContext) Gateways() context.ApplicationContextGateways {
	return &appContext.GatewaySet
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
