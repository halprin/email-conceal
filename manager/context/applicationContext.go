package context

type ApplicationContext interface {
	//controllers
	ConcealEmailController(arguments map[string]interface{}) (int, map[string]string)
	DeleteConcealEmailController(arguments map[string]interface{}) (int, map[string]string)

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
