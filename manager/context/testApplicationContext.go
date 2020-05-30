package context

type TestApplicationContext struct {
	ReceivedConcealEmailControllerArguments map[string]interface{}
	ReturnStatusFromConcealEmailController  int
	ReturnBodyFromConcealEmailController    map[string]string

	ReceivedDeleteConcealEmailControllerArguments map[string]interface{}
	ReturnStatusFromDeleteConcealEmailController  int
	ReturnBodyFromDeleteConcealEmailController    map[string]string

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

func (appContext *TestApplicationContext) ConcealEmailController(arguments map[string]interface{}) (int, map[string]string) {
	appContext.ReceivedConcealEmailControllerArguments = arguments
	return appContext.ReturnStatusFromConcealEmailController, appContext.ReturnBodyFromConcealEmailController
}

func (appContext *TestApplicationContext) DeleteConcealEmailController(arguments map[string]interface{}) (int, map[string]string) {
	appContext.ReceivedDeleteConcealEmailControllerArguments = arguments
	return appContext.ReturnStatusFromDeleteConcealEmailController, appContext.ReturnBodyFromDeleteConcealEmailController
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
