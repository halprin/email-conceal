package cli

import (
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/controllers/forwardEmail"
	"os"
)

var applicationContext = context.ApplicationContext{}

func Cli() {
	var forwardEmailController forwardEmail.ForwardEmail
	applicationContext.Resolve(&forwardEmailController)

	arguments := map[string]interface{}{
		"url": os.Args[1],
	}
	err := forwardEmailController.ForwardEmail(arguments)

	if err != nil {
		defer os.Exit(1) //allows for any other deferred actions before Exit is called
	}
}
