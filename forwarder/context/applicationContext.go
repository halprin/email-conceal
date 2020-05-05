package context

type ApplicationContext interface {
	//controllers
	ForwardEmailController(arguments map[string]interface{}) error

	//controller
	ReadEmailGateway(url string) ([]byte, error)

	SendEmailGateway(email []byte, recipients []string) error

	EnvironmentGateway(key string) string

	GetRealEmailForConcealPrefix(concealPrefix string) (string, error)

	//usecases
	ForwardEmailUsecase(url string) error

	//libraries
	Exit(returnCode int)
}
