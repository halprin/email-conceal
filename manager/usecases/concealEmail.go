package usecases

import (
	"fmt"
	"github.com/halprin/email-conceal/context"
	"github.com/halprin/email-conceal/entities"
)

func ConcealEmail(sourceEmail string, applicationContext context.ApplicationContext) (string, error) {
	err := entities.ValidateEmail(sourceEmail)
	if err != nil {
		return "", err
	}

	concealedEmailPrefix := applicationContext.GenerateRandomUuid()
	return fmt.Sprintf("%s@asdf.net", concealedEmailPrefix), nil
}
