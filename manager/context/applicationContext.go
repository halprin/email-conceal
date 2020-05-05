package context

type ApplicationContext interface {
	//controller
	ConcealEmailController(arguments []string) string

	//gateways
	EnvironmentGateway(key string) string

	//usecases
	AddConcealEmailUsecase(email string) (string, error)

	//libraries
	GenerateRandomUuid() string
	Exit(returnCode int)
}
