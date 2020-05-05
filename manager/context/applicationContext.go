package context

type ApplicationContext interface {
	//controller
	ConcealEmailController(arguments []string) string

	//usecases
	ConcealEmailUsecase(email string) (string, error)

	//libraries
	GenerateRandomUuid() string
	Exit(returnCode int)
}
