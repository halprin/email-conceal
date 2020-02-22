package context

type ApplicationContext interface {
	//gateways
	ForwardEmailGateway(arguments []string) error
	ReadEmailGateway(url string) ([]byte, error)
	SendEmailGateway(email []byte) error
	EnvironmentGateway(key string) string

	//usecases
	ForwardEmailUsecase(url string) error

	//libraries
	Exit(returnCode int)
}
