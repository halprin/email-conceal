package context

type ApplicationContext interface {
	//controller
	ConcealEmailController(arguments []string) string

	//gateways
	EnvironmentGateway(key string) string
	AddConcealedEmailToActualEmailMappingGateway(concealPrefix string, actualEmail string) error
	DeleteConcealedEmailToActualEmailMappingGateway(concealPrefix string) error

	//usecases
	AddConcealEmailUsecase(email string) (string, error)
	DeleteConcealEmailUsecase(concealPrefix string) error

	//libraries
	GenerateRandomUuid() string
	Exit(returnCode int)
}
