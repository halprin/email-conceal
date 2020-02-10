package usecases

import (
	"fmt"
	"github.com/halprin/email-conceal/forwarder/context"
	"github.com/halprin/email-conceal/forwarder/external/lib/errors"
)

func ForwardEmailUsecase(url string, applicationContext context.ApplicationContext) error {
	rawEmail, err := applicationContext.ReadEmailGateway(url)
	if err != nil {
		errorMessage := fmt.Sprintf("Unable to read e-mail at %s", url)
		return errors.Wrap(err, errorMessage)
	}

	err = applicationContext.SendEmailGateway(rawEmail)
	if err != nil {
		return errors.Wrap(err, "Unable to send e-mail")
	}

	return nil
}
