package actualEmail

import "github.com/halprin/email-conceal/src/entities"

type ActualEmailUsecase interface {
	Add(actualEmail string) error
}

type ActualEmailUsecaseImpl struct {}

func (receiver ActualEmailUsecaseImpl) Add(actualEmail string) error {
	err := entities.ValidateEmail(actualEmail)
	if err != nil {
		return err
	}

	//TODO: generate the secret
	//TODO: write actual e-mail with secret to database
	//TODO: send registration e-mail

	return nil
}
