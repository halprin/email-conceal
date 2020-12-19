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
	_ = forwardEmailController.ForwardEmail(arguments)
}
