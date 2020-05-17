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

	err = applicationContext.AddConcealEmailMappingGateway(concealedEmailPrefix, sourceEmail)
	if err != nil {
		return "", err
	}

	domain := applicationContext.EnvironmentGateway("DOMAIN")
	return fmt.Sprintf("%s@%s", concealedEmailPrefix, domain), nil
}
