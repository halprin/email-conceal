package gateways

import (
	"fmt"
	"github.com/halprin/email-conceal/forwarder/context"
)

func CliForwardEmail(cliArguments []string, applicationContext context.ApplicationContext) error {
	url := cliArguments[1]
	fmt.Println("URL to read e-mail from =", url)

	err := applicationContext.ForwardEmailUsecase(url)
	if err != nil {
		fmt.Println("Unable to forward e-mail")
		defer applicationContext.Exit(1) //allows for any other deferred actions before Exit is called
		return err
	}

	fmt.Println("Forwarded e-mail")
	return nil
}