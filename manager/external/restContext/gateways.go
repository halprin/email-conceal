package restContext

import (
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/gateways"
)

type RestApplicationContextGateways struct{
	ParentContext context.ApplicationContext
}

func (appContextGateways *RestApplicationContextGateways) EnvironmentGateway(key string) string {
	return gateways.OsEnvEnvironmentGateway(key, appContextGateways.ParentContext)
}

func (appContextGateways *RestApplicationContextGateways) AddConcealedEmailToActualEmailMappingGateway(concealPrefix string, actualEmail string) error {
	return gateways.AddConcealedEmailToActualEmailMapping(concealPrefix, actualEmail, appContextGateways.ParentContext)
}

func (appContextGateways *RestApplicationContextGateways) DeleteConcealedEmailToActualEmailMappingGateway(concealPrefix string) error {
	return gateways.DeleteConcealedEmailToActualEmailMapping(concealPrefix, appContextGateways.ParentContext)
}