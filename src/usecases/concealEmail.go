package usecases

import (
	"fmt"
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/entities"
	"github.com/halprin/email-conceal/manager/external/lib/errors"
)


var applicationContext = context.ApplicationContext{}

type ConcealEmailUsecase interface {
	Add(sourceEmail string, description *string) (string, error)
	Delete(concealedEmailPrefix string) error
	AddDescriptionToExistingEmail(concealedEmailPrefix string, description string) error
	DeleteDescriptionFromExistingEmail(concealedEmailPrefix string) error
}

type ConcealEmailNotExistError struct {
	ConcealEmailId string
}

func (c ConcealEmailNotExistError) Error() string {
	return fmt.Sprintf("The conceal e-mail %s doesn't exist", c.ConcealEmailId)
}

type ConcealEmailUsecaseImpl struct {}

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

	var uuidLibrary context.UuidLibrary
	applicationContext.Resolve(&uuidLibrary)
	concealedEmailPrefix := uuidLibrary.GenerateRandomUuid()

	var concealEmailGateway ConcealEmailGateway
	applicationContext.Resolve(&concealEmailGateway)
	err = concealEmailGateway.AddConcealedEmailToActualEmailMapping(concealedEmailPrefix, sourceEmail, description)
	if err != nil {
		return "", errors.Wrap(err, "Unable to add conceal e-mail to actual e-mail mapping")
	}

	var environmentGateway context.EnvironmentGateway
	applicationContext.Resolve(&environmentGateway)
	domain := environmentGateway.GetEnvironmentValue("DOMAIN")

	return fmt.Sprintf("%s@%s", concealedEmailPrefix, domain), nil
}

func (receiver ConcealEmailUsecaseImpl) Delete(concealedEmailPrefix string) error {
	var concealEmailGateway ConcealEmailGateway
	applicationContext.Resolve(&concealEmailGateway)
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

	var concealEmailGateway ConcealEmailGateway
	applicationContext.Resolve(&concealEmailGateway)
	err = concealEmailGateway.UpdateConcealedEmail(concealedEmailPrefix, &description)
	if err != nil {
		return errors.Wrap(err, "Unable to update description of conceal e-mail")
	}

	return nil
}

func (receiver ConcealEmailUsecaseImpl) DeleteDescriptionFromExistingEmail(concealedEmailPrefix string) error {

	var concealEmailGateway ConcealEmailGateway
	applicationContext.Resolve(&concealEmailGateway)
	err := concealEmailGateway.UpdateConcealedEmail(concealedEmailPrefix, nil)
	if err != nil {
		return errors.Wrap(err, "Unable to delete description of conceal e-mail")
	}

	return nil
}
