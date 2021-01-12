package actualEmail


type SendRegistrationEmailGateway interface {
	SendEmail(email []byte, recipient string) error
}

type ActualEmailConfigurationGateway interface {
	AddUnprovedActualEmail(actualEmail string, ownershipSecret string) error
}

