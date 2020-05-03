package context

type TestApplicationContext struct {
	ReceivedForwardEmailControllerArguments map[string]interface{}
	ReturnErrorFromForwardEmailController   error

	ReceivedReadEmailGatewayArguments string
	ReturnFromReadEmailGateway        []byte
	ReturnErrorFromReadEmailGateway   error

	ReceivedSendEmailGatewayEmailArgument []byte
	ReceivedSendEmailGatewayRecipientArgument []string
	ReturnErrorFromSendEmailGateway       error

	ReceivedEnvironmentGatewayArguments string
	ReturnFromEnvironmentGateway        string

	ReceivedGetRealEmailForConcealPrefixArguments string
	ReturnFromGetRealEmailForConcealPrefix        string
	ReturnErrorFromGetRealEmailForConcealPrefix   error

	ReceivedForwardEmailUsecaseArguments string
	ReturnErrorForwardEmailUsecase       error

	ReceivedExitReturnCode                    int
}

func (appContext *TestApplicationContext) ForwardEmailController(arguments map[string]interface{}) error {
	appContext.ReceivedForwardEmailControllerArguments = arguments
	return appContext.ReturnErrorFromForwardEmailController
}

func (appContext *TestApplicationContext) ReadEmailGateway(url string) ([]byte, error) {
	appContext.ReceivedReadEmailGatewayArguments = url
	return appContext.ReturnFromReadEmailGateway, appContext.ReturnErrorFromReadEmailGateway
}

func (appContext *TestApplicationContext) SendEmailGateway(email []byte, recipients []string) error {
	appContext.ReceivedSendEmailGatewayEmailArgument = email
	appContext.ReceivedSendEmailGatewayRecipientArgument = recipients
	return appContext.ReturnErrorFromSendEmailGateway
}

func (appContext *TestApplicationContext) EnvironmentGateway(key string) string {
	appContext.ReceivedEnvironmentGatewayArguments = key
	return appContext.ReturnFromEnvironmentGateway
}

func (appContext *TestApplicationContext) GetRealEmailForConcealPrefix(concealPrefix string) (string, error) {
	appContext.ReceivedGetRealEmailForConcealPrefixArguments = concealPrefix
	return appContext.ReturnFromGetRealEmailForConcealPrefix, appContext.ReturnErrorFromGetRealEmailForConcealPrefix
}

func (appContext *TestApplicationContext) ForwardEmailUsecase(url string) error {
	appContext.ReceivedForwardEmailUsecaseArguments = url
	return appContext.ReturnErrorForwardEmailUsecase
}

func (appContext *TestApplicationContext) Exit(returnCode int) {
	appContext.ReceivedExitReturnCode = returnCode
}
