package forwardEmail

import (
	"fmt"
	forwardEmailUsecase "github.com/halprin/email-conceal/src/usecases/forwardEmail"
	"os"
)

type CliForwardController struct {}

func (receiver CliForwardController) ForwardEmail(cliArguments map[string]interface{}) error {
	url := cliArguments["url"].(string)
	fmt.Println("URL to read e-mail from =", url)

	var forwardEmailUsecaseVar forwardEmailUsecase.ForwardEmailUsecase
	applicationContext.Resolve(&forwardEmailUsecaseVar)

	err := forwardEmailUsecaseVar.ForwardEmail(url)
	if err != nil {
		fmt.Println("Unable to forward e-mail")
		defer os.Exit(1) //allows for any other deferred actions before Exit is called
		return err
	}

	fmt.Println("Forwarded e-mail")
	return nil
}
