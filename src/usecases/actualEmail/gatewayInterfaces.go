package actualEmail

type ActualEmailConfigurationGateway interface {
	AddUnprovedActualEmail(actualEmail string, ownershipSecret string) error
}
