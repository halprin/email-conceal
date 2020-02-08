package context

type ApplicationContext interface {
	//gateways
	ConcealEmailGateway(cliArguments []string) string

	//usecases
	ConcealEmailUsecase(email string) string

	//libraries
	GenerateRandomUuid() string
}
