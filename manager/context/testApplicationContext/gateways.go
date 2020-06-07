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

func (appContextGateways *TestApplicationContextGateways) GetEnvironmentValue(key string) string {
	appContextGateways.ReceivedEnvironmentGatewayArguments = key
	return appContextGateways.ReturnFromEnvironmentGateway[key]
}

func (appContextGateways *TestApplicationContextGateways) AddConcealedEmailToActualEmailMapping(concealPrefix string, actualEmail string, description *string) error {
	appContextGateways.ReceivedAddConcealedEmailToActualEmailMappingGatewayConcealPrefixArgument = concealPrefix
	appContextGateways.ReceivedAddConcealedEmailToActualEmailMappingGatewayEmailArgument = actualEmail
	return appContextGateways.ReturnErrorFromAddConcealedEmailToActualEmailMappingGateway
}

func (appContextGateways *TestApplicationContextGateways) DeleteConcealedEmailToActualEmailMapping(concealPrefix string) error {
	appContextGateways.ReceivedDeleteConcealedEmailToActualEmailMappingGatewayConcealPrefixArgument = concealPrefix
	return appContextGateways.ReturnErrorFromDeleteConcealedEmailToActualEmailMappingGateway
}

