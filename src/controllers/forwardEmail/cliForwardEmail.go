package forwardEmail

import (
	"fmt"
	forwardEmailUsecase "github.com/halprin/email-conceal/src/usecases/forwardEmail"
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
		return err
	}

	fmt.Println("Forwarded e-mail")
	return nil
}
