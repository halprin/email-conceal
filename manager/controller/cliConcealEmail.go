package controller

import (
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/entities"
	"github.com/halprin/email-conceal/manager/external/lib/errors"
	"log"
)

func CliConcealEmailController(cliArguments []string, applicationContext context.ApplicationContext) string {
	sourceEmail := cliArguments[1]
	log.Println("E-mail to conceal =", sourceEmail)

	concealedEmail, err := applicationContext.AddConcealEmailUsecase(sourceEmail)
	if errors.Is(err, entities.InvalidEmailAddressError) {
		log.Printf("E-mail %s is invalid\n", sourceEmail)
		defer applicationContext.Exit(1) //allows for any other deferred actions before Exit is called
		return ""
	} else if err != nil {
		log.Printf("Another error occured, %+v", err)
		defer applicationContext.Exit(2)
		return ""
	}

	log.Println("Concealed e-mail address =", concealedEmail)
	return concealedEmail
}
