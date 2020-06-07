package testApplicationContext


type TestApplicationContextLibraries struct{
	ReturnFromGenerateRandomUuid string

	ReceivedExitReturnCode int
}

func (appContextLibraries *TestApplicationContextLibraries) GenerateRandomUuid() string {
	return appContextLibraries.ReturnFromGenerateRandomUuid
}

func (appContextLibraries *TestApplicationContextLibraries) Exit(returnCode int) {
	appContextLibraries.ReceivedExitReturnCode = returnCode
}
