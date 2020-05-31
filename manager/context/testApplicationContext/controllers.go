package testApplicationContext


type TestApplicationContextControllers struct{
	ReceivedConcealEmailControllerArguments map[string]interface{}
	ReturnStatusFromConcealEmailController  int
	ReturnBodyFromConcealEmailController    map[string]string

	ReceivedDeleteConcealEmailControllerArguments map[string]interface{}
	ReturnStatusFromDeleteConcealEmailController  int
	ReturnBodyFromDeleteConcealEmailController    map[string]string
}

func (appContextControllers *TestApplicationContextControllers) ConcealEmailController(arguments map[string]interface{}) (int, map[string]string) {
	appContextControllers.ReceivedConcealEmailControllerArguments = arguments
	return appContextControllers.ReturnStatusFromConcealEmailController, appContextControllers.ReturnBodyFromConcealEmailController
}

func (appContextControllers *TestApplicationContextControllers) DeleteConcealEmailController(arguments map[string]interface{}) (int, map[string]string) {
	appContextControllers.ReceivedDeleteConcealEmailControllerArguments = arguments
	return appContextControllers.ReturnStatusFromDeleteConcealEmailController, appContextControllers.ReturnBodyFromDeleteConcealEmailController
}

