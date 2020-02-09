package context

type TestApplicationContext struct {
	ReceivedConcealEmailGatewayArguments []string
	ReturnFromConcealEmailGateway        string

	ReceivedConcealEmailUsecaseEmail   string
	ReturnFromConcealEmailUsecase      string
	ReturnErrorFromConcealEmailUsecase error

	ReturnFromGenerateRandomUuid string

	ReceivedExitReturnCode int
}

func (appContext *TestApplicationContext) ConcealEmailGateway(arguments []string) string {
	appContext.ReceivedConcealEmailGatewayArguments = arguments
	return appContext.ReturnFromConcealEmailGateway
}

func (appContext *TestApplicationContext) ConcealEmailUsecase(email string) (string, error) {
	appContext.ReceivedConcealEmailUsecaseEmail = email
	return appContext.ReturnFromConcealEmailUsecase, appContext.ReturnErrorFromConcealEmailUsecase
}

func (appContext *TestApplicationContext) GenerateRandomUuid() string {
	return appContext.ReturnFromGenerateRandomUuid
}

func (appContext *TestApplicationContext) Exit(returnCode int) {
	appContext.ReceivedExitReturnCode = returnCode
}