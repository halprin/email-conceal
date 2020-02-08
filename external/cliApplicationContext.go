package external

import (
	"github.com/halprin/email-conceal/external/lib"
	"github.com/halprin/email-conceal/gateways"
	"github.com/halprin/email-conceal/usecases"
	"os"
)

type CliApplicationContext struct{}

func (cliAppContext CliApplicationContext) ConcealEmailGateway(cliArguments []string) string {
	return gateways.CliConcealEmailGateway(cliArguments, cliAppContext)
}

func (cliAppContext CliApplicationContext) ConcealEmailUsecase(email string) (string, error) {
	return usecases.ConcealEmail(email, cliAppContext)
}

func (cliAppContext CliApplicationContext) GenerateRandomUuid() string {
	return lib.GenerateGoogleRandomUuid(cliAppContext)
}

func (cliAppContext CliApplicationContext) Exit(returnCode int) {
	os.Exit(returnCode)
}
