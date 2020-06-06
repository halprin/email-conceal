package context

type ApplicationContext interface {
	//controllers
	Controllers() ApplicationContextControllers

	//gateways
	Gateways() ApplicationContextGateways

	//usecases
	Usecases() ApplicationContextUsecases

	//libraries
	GenerateRandomUuid() string
	Exit(returnCode int)
}
