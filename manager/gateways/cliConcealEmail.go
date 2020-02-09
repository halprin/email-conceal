package gateways

import (
	"fmt"
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/entities"
	"github.com/halprin/email-conceal/manager/external/lib/errors"
)

func CliConcealEmailGateway(cliArguments []string, applicationContext context.ApplicationContext) string {
	sourceEmail := cliArguments[1]
	fmt.Println("E-mail to conceal =", sourceEmail)

	concealedEmail, err := applicationContext.ConcealEmailUsecase(sourceEmail)
	if errors.Is(err, entities.InvalidEmailAddressError) {
		fmt.Printf("e-mail %s is invalid\n", sourceEmail)
		defer applicationContext.Exit(1) //allows for any other deferred actions before Exit is called
		return ""
	}

	fmt.Println("Concealed e-mail address =", concealedEmail)
	return concealedEmail
}
