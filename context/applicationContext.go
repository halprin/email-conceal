package context

type ApplicationContext interface {
	//gateways
	ConcealEmailGateway(arguments []string) string

	//usecases
	ConcealEmailUsecase(email string) (string, error)

	//libraries
	GenerateRandomUuid() string
	Exit(returnCode int)
}
