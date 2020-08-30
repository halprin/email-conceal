package restContext

import (
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/gateways"
)

type RestApplicationContextGateways struct{
	ParentContext context.ApplicationContext
}

func (appContextGateways *RestApplicationContextGateways) GetEnvironmentValue(key string) string {
	return gateways.OsEnvEnvironmentGateway(key, appContextGateways.ParentContext)
}

func (appContextGateways *RestApplicationContextGateways) AddConcealedEmailToActualEmailMapping(concealPrefix string, actualEmail string, description *string) error {
	return gateways.AddConcealedEmailToActualEmailMapping(concealPrefix, actualEmail, description, appContextGateways.ParentContext)
}

func (appContextGateways *RestApplicationContextGateways) DeleteConcealedEmailToActualEmailMapping(concealPrefix string) error {
	return gateways.DeleteConcealedEmailToActualEmailMapping(concealPrefix, appContextGateways.ParentContext)
}
