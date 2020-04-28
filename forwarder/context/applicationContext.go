package context

type ApplicationContext interface {
	//gateways
	ForwardEmailController(arguments map[string]interface{}) error
	ReadEmailGateway(url string) ([]byte, error)
	SendEmailGateway(email []byte) error
	EnvironmentGateway(key string) string

	//usecases
	ForwardEmailUsecase(url string) error

	//libraries
	Exit(returnCode int)
}
