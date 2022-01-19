package actualEmail

type ActualEmailConfigurationGateway interface {
	AddUnprovedActualEmail(actualEmail string, ownershipSecret string) error
	GetActualEmailForSecret(secret string) (*string, error)
	ActivateActualEmail(actualEmail string) error
}
