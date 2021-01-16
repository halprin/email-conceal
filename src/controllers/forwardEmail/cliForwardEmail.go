package forwardEmail

import (
	forwardEmailUsecase "github.com/halprin/email-conceal/src/usecases/forwardEmail"
	"log"
)

type CliForwardController struct {}

func (receiver CliForwardController) ForwardEmail(cliArguments map[string]interface{}) error {
	url := cliArguments["url"].(string)
	log.Println("URL to read e-mail from =", url)

	var forwardEmailUsecaseVar forwardEmailUsecase.ForwardEmailUsecase
	applicationContext.Resolve(&forwardEmailUsecaseVar)

	err := forwardEmailUsecaseVar.ForwardEmail(url)
	if err != nil {
		log.Println("Unable to forward e-mail")
		return err
	}

	return nil
}
