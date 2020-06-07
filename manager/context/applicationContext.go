package context

type ApplicationContext interface {
	//controllers
	Controllers() ApplicationContextControllers

	//gateways
	Gateways() ApplicationContextGateways

	//usecases
	Usecases() ApplicationContextUsecases

	//libraries
	Libraries() ApplicationContextLibraries
}
