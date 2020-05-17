package context

type ApplicationContext interface {
	//controller
	ConcealEmailController(arguments []string) string

	//gateways
	EnvironmentGateway(key string) string
	AddConcealEmailMappingGateway(concealPrefix string, actualEmail string) error
	DeleteConcealEmailMappingGateway(concealPrefix string, actualEmail string) error

	//usecases
	AddConcealEmailUsecase(email string) (string, error)

	//libraries
	GenerateRandomUuid() string
	Exit(returnCode int)
}
