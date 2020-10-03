package usecases

import (
	"fmt"
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/entities"
	"github.com/halprin/email-conceal/manager/external/lib/errors"
)


type ConcealEmailUsecase struct {}

func (receiver ConcealEmailUsecase) Add(sourceEmail string, description *string) (string, error) {
	err := entities.ValidateEmail(sourceEmail)
	if err != nil {
		return "", err
	}

	if description != nil {
		//the description exists, so validate it
		err := entities.ValidateDescription(*description)
		if err != nil {
			return "", err
		}
	}

	concealedEmailPrefix := applicationContext.Libraries().GenerateRandomUuid()
	err = applicationContext.Gateways().AddConcealedEmailToActualEmailMapping(concealedEmailPrefix, sourceEmail, description)
	if err != nil {
		return "", errors.Wrap(err, "Unable to add conceal e-mail to actual e-mail mapping")
	}

	domain := applicationContext.Gateways().GetEnvironmentValue("DOMAIN")

	return fmt.Sprintf("%s@%s", concealedEmailPrefix, domain), nil
}

func (receiver ConcealEmailUsecase) Delete(concealedEmailPrefix string) error {
	err := applicationContext.Gateways().DeleteConcealedEmailToActualEmailMapping(concealedEmailPrefix)
	if err != nil {
		return errors.Wrap(err, "Unable to delete conceal e-mail to actual e-mail mapping")
	}

	return nil
}
