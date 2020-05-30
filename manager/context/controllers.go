package context

type ApplicationContextControllers interface {
	ConcealEmailController(arguments map[string]interface{}) (int, map[string]string)
	DeleteConcealEmailController(arguments map[string]interface{}) (int, map[string]string)
}
