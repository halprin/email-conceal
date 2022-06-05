package account

import (
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/entities"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"github.com/halprin/email-conceal/src/usecases/actualEmail"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var applicationContext = context.ApplicationContext{}
var accountConfigurationGateway AccountConfigurationGateway
var actualEmailUsecase actualEmail.ActualEmailUsecase

type AccountUsecase interface {
	Create(emailUsername string, password string) error
}

type AccountUsecaseImpl struct {
}

func (receiver AccountUsecaseImpl) Init() {
	applicationContext.Resolve(&accountConfigurationGateway)
	applicationContext.Resolve(&actualEmailUsecase)
}

func (receiver AccountUsecaseImpl) Create(emailUsername string, password string) error {
	log.Println("Creating a new account")

	err := entities.ValidateEmail(emailUsername)
	if err != nil {
		return err
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return errors.Wrap(err, "Unable to hash your password")
	}

	err = accountConfigurationGateway.AddAccount(emailUsername, hashedPassword)
	if err != nil {
		return errors.Wrap(err, "Unable to add the new user account")
	}

	err = actualEmailUsecase.Add(emailUsername)
	if err != nil {
		return errors.Wrap(err, "Wasn't able to register the e-mail address to forward e-mails to")
	}

	return nil
}

func hashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	cipherText, err := bcrypt.GenerateFromPassword(passwordBytes, 13)
	if err != nil {
		return "", err
	}

	return string(cipherText), nil
}
