package context

type TestApplicationContext struct {
	ReceivedForwardEmailGatewayArguments map[string]interface{}
	ReturnErrorFromForwardEmailGateway   error

	ReceivedReadEmailGatewayArguments string
	ReturnFromReadEmailGateway        []byte
	ReturnErrorFromReadEmailGateway   error

	ReceivedSendEmailGatewayArguments []byte
	ReturnErrorFromSendEmailGateway   error

	ReceivedEnvironmentGatewayArguments string
	ReturnFromEnvironmentGateway        string

	ReceivedForwardEmailUsecaseArguments string
	ReturnErrorForwardEmailUsecase       error

	ReceivedExitReturnCode int
}

func (appContext *TestApplicationContext) ForwardEmailController(arguments map[string]interface{}) error {
	appContext.ReceivedForwardEmailGatewayArguments = arguments
	return appContext.ReturnErrorFromForwardEmailGateway
}

func (appContext *TestApplicationContext) ReadEmailGateway(url string) ([]byte, error) {
	appContext.ReceivedReadEmailGatewayArguments = url
	return appContext.ReturnFromReadEmailGateway, appContext.ReturnErrorFromReadEmailGateway
}

func (appContext *TestApplicationContext) SendEmailGateway(email []byte) error {
	appContext.ReceivedSendEmailGatewayArguments = email
	return appContext.ReturnErrorFromSendEmailGateway
}

func (appContext *TestApplicationContext) EnvironmentGateway(key string) string {
	appContext.ReceivedEnvironmentGatewayArguments = key
	return appContext.ReturnFromEnvironmentGateway
}

func (appContext *TestApplicationContext) ForwardEmailUsecase(url string) error {
	appContext.ReceivedForwardEmailUsecaseArguments = url
	return appContext.ReturnErrorForwardEmailUsecase
}

func (appContext *TestApplicationContext) Exit(returnCode int) {
	appContext.ReceivedExitReturnCode = returnCode
}
