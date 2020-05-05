package usecases

import (
	"fmt"
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/entities"
)

func AddConcealEmailUsecase(sourceEmail string, applicationContext context.ApplicationContext) (string, error) {
	err := entities.ValidateEmail(sourceEmail)
	if err != nil {
		return "", err
	}

	concealedEmailPrefix := applicationContext.GenerateRandomUuid()
	return fmt.Sprintf("%s@asdf.net", concealedEmailPrefix), nil
}
