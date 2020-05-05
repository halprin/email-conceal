package context

type TestApplicationContext struct {
	ReceivedConcealEmailControllerArguments []string
	ReturnFromConcealEmailController        string

	ReceivedConcealEmailUsecaseEmail   string
	ReturnFromConcealEmailUsecase      string
	ReturnErrorFromConcealEmailUsecase error

	ReturnFromGenerateRandomUuid string

	ReceivedExitReturnCode int
}

func (appContext *TestApplicationContext) ConcealEmailController(arguments []string) string {
	appContext.ReceivedConcealEmailControllerArguments = arguments
	return appContext.ReturnFromConcealEmailController
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
