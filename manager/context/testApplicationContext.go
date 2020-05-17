package context

type TestApplicationContext struct {
	ReceivedConcealEmailControllerArguments []string
	ReturnFromConcealEmailController        string

	ReceivedEnvironmentGatewayArguments string
	ReturnFromEnvironmentGateway        map[string]string

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

func (appContext *TestApplicationContext) EnvironmentGateway(key string) string {
	appContext.ReceivedEnvironmentGatewayArguments = key
	return appContext.ReturnFromEnvironmentGateway[key]
}

func (appContext *TestApplicationContext) AddConcealedEmailToActualEmailMappingGateway(concealPrefix string, actualEmail string) error {
	//TODO: fill in
	return nil
}

func (appContext *TestApplicationContext) DeleteConcealedEmailToActualEmailMappingGateway(concealPrefix string) error {
	//TODO: fill in
	return nil
}

func (appContext *TestApplicationContext) AddConcealEmailUsecase(email string) (string, error) {
	appContext.ReceivedConcealEmailUsecaseEmail = email
	return appContext.ReturnFromConcealEmailUsecase, appContext.ReturnErrorFromConcealEmailUsecase
}

func (appContext *TestApplicationContext) DeleteConcealEmailUsecase(concealPrefix string) error {
	//TODO: fill in
	return nil
}

func (appContext *TestApplicationContext) GenerateRandomUuid() string {
	return appContext.ReturnFromGenerateRandomUuid
}

func (appContext *TestApplicationContext) Exit(returnCode int) {
	appContext.ReceivedExitReturnCode = returnCode
}
