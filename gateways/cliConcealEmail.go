package gateways

import (
	"email-conceal/context"
	"fmt"
)

func CliConcealEmailGateway(cliArguments []string, applicationContext context.ApplicationContext) string {
	sourceEmail := cliArguments[1]
	fmt.Println("E-mail to conceal =", sourceEmail)

	return applicationContext.ConcealEmailUsecase(sourceEmail)
}