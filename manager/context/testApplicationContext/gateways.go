package testApplicationContext


type TestApplicationContextGateways struct{
	ReceivedEnvironmentGatewayArguments string
	ReturnFromEnvironmentGateway        map[string]string

	ReceivedAddConcealedEmailToActualEmailMappingGatewayConcealPrefixArgument string
	ReceivedAddConcealedEmailToActualEmailMappingGatewayEmailArgument         string
	ReturnErrorFromAddConcealedEmailToActualEmailMappingGateway               error

	ReceivedDeleteConcealedEmailToActualEmailMappingGatewayConcealPrefixArgument string
	ReturnErrorFromDeleteConcealedEmailToActualEmailMappingGateway               error
}

func (appContextGateways *TestApplicationContextGateways) EnvironmentGateway(key string) string {
	appContextGateways.ReceivedEnvironmentGatewayArguments = key
	return appContextGateways.ReturnFromEnvironmentGateway[key]
}

func (appContextGateways *TestApplicationContextGateways) AddConcealedEmailToActualEmailMappingGateway(concealPrefix string, actualEmail string) error {
	appContextGateways.ReceivedAddConcealedEmailToActualEmailMappingGatewayConcealPrefixArgument = concealPrefix
	appContextGateways.ReceivedAddConcealedEmailToActualEmailMappingGatewayEmailArgument = actualEmail
	return appContextGateways.ReturnErrorFromAddConcealedEmailToActualEmailMappingGateway
}

func (appContextGateways *TestApplicationContextGateways) DeleteConcealedEmailToActualEmailMappingGateway(concealPrefix string) error {
	appContextGateways.ReceivedDeleteConcealedEmailToActualEmailMappingGatewayConcealPrefixArgument = concealPrefix
	return appContextGateways.ReturnErrorFromDeleteConcealedEmailToActualEmailMappingGateway
}

