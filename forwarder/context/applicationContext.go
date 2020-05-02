package context

type ApplicationContext interface {
	//controllers
	ForwardEmailController(arguments map[string]interface{}) error

	//gateways
	ReadEmailGateway(url string) ([]byte, error)

	SendEmailGateway(email []byte, recipient string) error

	EnvironmentGateway(key string) string

	GetRealEmailForConcealPrefix(concealPrefix string) (string, error)

	//usecases
	ForwardEmailUsecase(url string) error

	//libraries
	Exit(returnCode int)
}
