package forwardEmail

import (
	"fmt"
)

type CliForwardController struct {}

func (receiver CliForwardController) ForwardEmail(cliArguments map[string]interface{}) error {
	url := cliArguments["url"].(string)
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
