package context

type ApplicationContextGateways interface {
	GetEnvironmentValue(key string) string
	AddConcealedEmailToActualEmailMapping(concealPrefix string, actualEmail string, description *string) error
	DeleteConcealedEmailToActualEmailMapping(concealPrefix string) error
}
