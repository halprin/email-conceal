package actualEmail

//TODO: need to implement
type SendRegistrationEmailGateway interface {
	SendEmail(email []byte, recipient string) error
}

//TODO: need to implement
type ActualEmailConfigurationGateway interface {
	AddUnprovedActualEmail(actualEmail string, ownershipSecret string) error
}
