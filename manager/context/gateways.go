package context

type ApplicationContextGateways interface {
	EnvironmentGateway(key string) string
	AddConcealedEmailToActualEmailMappingGateway(concealPrefix string, actualEmail string) error
	DeleteConcealedEmailToActualEmailMappingGateway(concealPrefix string) error
}
