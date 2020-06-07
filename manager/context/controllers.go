package context

type ApplicationContextControllers interface {
	ConcealEmail(arguments map[string]interface{}) (int, map[string]string)
	DeleteConcealEmail(arguments map[string]interface{}) (int, map[string]string)
}
