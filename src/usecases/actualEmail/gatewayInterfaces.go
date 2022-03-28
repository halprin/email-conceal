package actualEmail

import "github.com/halprin/email-conceal/src/external/lib/errors"

type ActualEmailConfigurationGateway interface {
	AddUnprovedActualEmail(actualEmail string, ownershipSecret string) error
	GetActualEmailForSecret(secret string) (*string, error)
	ActivateActualEmail(actualEmail string) error
}

var ActualEmailDoesNotExist = errors.New("Actual e-mail does not exist")
