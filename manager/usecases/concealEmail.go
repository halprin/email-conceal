package usecases

import (
	"fmt"
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/entities"
	"github.com/halprin/email-conceal/manager/external/lib/errors"
)


var applicationContext = context.ApplicationContext{}
var concealEmailGateway ConcealEmailGateway
var environmentGateway context.EnvironmentGateway
var uuidLibrary context.UuidLibrary

func init() {
	applicationContext.Resolve(&concealEmailGateway)
	applicationContext.Resolve(&environmentGateway)
	applicationContext.Resolve(&uuidLibrary)
}

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

	concealedEmailPrefix := uuidLibrary.GenerateRandomUuid()
	err = concealEmailGateway.AddConcealedEmailToActualEmailMapping(concealedEmailPrefix, sourceEmail, description)
	if err != nil {
		return "", errors.Wrap(err, "Unable to add conceal e-mail to actual e-mail mapping")
	}

	domain := environmentGateway.GetEnvironmentValue("DOMAIN")

	return fmt.Sprintf("%s@%s", concealedEmailPrefix, domain), nil
}

func (receiver ConcealEmailUsecase) Delete(concealedEmailPrefix string) error {
	err := concealEmailGateway.DeleteConcealedEmailToActualEmailMapping(concealedEmailPrefix)
	if err != nil {
		return errors.Wrap(err, "Unable to delete conceal e-mail to actual e-mail mapping")
	}

	return nil
}
