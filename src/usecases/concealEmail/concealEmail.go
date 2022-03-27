package concealEmail

import (
	"fmt"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/entities"
	"github.com/halprin/email-conceal/src/external/lib/errors"
)

var applicationContext = context.ApplicationContext{}
var uuidLibrary context.UuidLibrary
var concealEmailGateway ConcealEmailGateway
var environmentGateway context.EnvironmentGateway

type ConcealEmailUsecase interface {
	Add(sourceEmail string, description *string) (string, error)
	Delete(concealedEmailPrefix string) error
	AddDescriptionToExistingEmail(concealedEmailPrefix string, description string) error
	DeleteDescriptionFromExistingEmail(concealedEmailPrefix string) error
}

type ConcealEmailUsecaseImpl struct{}

func (receiver ConcealEmailUsecaseImpl) Init() {
	applicationContext.Resolve(&uuidLibrary)
	applicationContext.Resolve(&concealEmailGateway)
	applicationContext.Resolve(&environmentGateway)
}

func (receiver ConcealEmailUsecaseImpl) Add(sourceEmail string, description *string) (string, error) {
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

	emailIsVerified, err := receiver.actualEmailIsVerified(sourceEmail)
	if err != nil {
		return "", errors.Wrap(err, "Unable to determine e-mail ownership due to error")
	} else if !emailIsVerified {
		return "", errors.New("Provided e-mail ownership has not been verified")
	}

	concealedEmailPrefix := uuidLibrary.GenerateRandomUuid()

	err = concealEmailGateway.AddConcealedEmailToActualEmailMapping(concealedEmailPrefix, sourceEmail, description)
	if err != nil {
		return "", errors.Wrap(err, "Unable to add conceal e-mail to actual e-mail mapping")
	}

	domain := environmentGateway.GetEnvironmentValue("DOMAIN")

	return fmt.Sprintf("%s@%s", concealedEmailPrefix, domain), nil
}

func (receiver ConcealEmailUsecaseImpl) actualEmailIsVerified(sourceEmail string) (bool, error) {
	return false, nil
}

func (receiver ConcealEmailUsecaseImpl) Delete(concealedEmailPrefix string) error {
	err := concealEmailGateway.DeleteConcealedEmailToActualEmailMapping(concealedEmailPrefix)
	if err != nil {
		return errors.Wrap(err, "Unable to delete conceal e-mail to actual e-mail mapping")
	}

	return nil
}

func (receiver ConcealEmailUsecaseImpl) AddDescriptionToExistingEmail(concealedEmailPrefix string, description string) error {
	err := entities.ValidateDescription(description)
	if err != nil {
		return err
	}

	err = concealEmailGateway.UpdateConcealedEmail(concealedEmailPrefix, &description)
	if err != nil {
		return errors.Wrap(err, "Unable to update description of conceal e-mail")
	}

	return nil
}

func (receiver ConcealEmailUsecaseImpl) DeleteDescriptionFromExistingEmail(concealedEmailPrefix string) error {
	err := concealEmailGateway.UpdateConcealedEmail(concealedEmailPrefix, nil)
	if err != nil {
		return errors.Wrap(err, "Unable to delete description of conceal e-mail")
	}

	return nil
}
