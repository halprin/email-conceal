package context

type ApplicationContextLibraries interface {
	GenerateRandomUuid() string
	Exit(returnCode int)
}
