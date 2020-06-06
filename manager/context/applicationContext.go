package context

type ApplicationContext interface {
	//controllers
	Controllers() ApplicationContextControllers

	//gateways
	Gateways() ApplicationContextGateways

	//usecases
	AddConcealEmailUsecase(email string) (string, error)
	DeleteConcealEmailUsecase(concealPrefix string) error

	//libraries
	GenerateRandomUuid() string
	Exit(returnCode int)
}
