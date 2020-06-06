package testApplicationContext

import "github.com/halprin/email-conceal/manager/context"

type TestApplicationContext struct {
	ControllerSet TestApplicationContextControllers
	GatewaySet    TestApplicationContextGateways
	UsecaseSet    TestApplicationContextUsecases

	ReturnFromGenerateRandomUuid string

	ReceivedExitReturnCode int
}

func (appContext *TestApplicationContext) Controllers() context.ApplicationContextControllers {
	return &appContext.ControllerSet
}

func (appContext *TestApplicationContext) Gateways() context.ApplicationContextGateways {
	return &appContext.GatewaySet
}

func (appContext *TestApplicationContext) Usecases() context.ApplicationContextUsecases {
	return &appContext.UsecaseSet
}

func (appContext *TestApplicationContext) TestUsecases() TestApplicationContextUsecases {
	return appContext.UsecaseSet
}

func (appContext *TestApplicationContext) GenerateRandomUuid() string {
	return appContext.ReturnFromGenerateRandomUuid
}

func (appContext *TestApplicationContext) Exit(returnCode int) {
	appContext.ReceivedExitReturnCode = returnCode
}
