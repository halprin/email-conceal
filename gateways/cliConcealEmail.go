package gateways

import (
	"fmt"
	"github.com/halprin/email-conceal/context"
)

func CliConcealEmailGateway(cliArguments []string, applicationContext context.ApplicationContext) string {
	sourceEmail := cliArguments[1]
	fmt.Println("E-mail to conceal =", sourceEmail)

	return applicationContext.ConcealEmailUsecase(sourceEmail)
}
