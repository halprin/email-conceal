package context

type ApplicationContextGateways interface {
	GetEnvironmentValue(key string) string
	AddConcealedEmailToActualEmailMapping(concealPrefix string, actualEmail string) error
	DeleteConcealedEmailToActualEmailMapping(concealPrefix string) error
}
