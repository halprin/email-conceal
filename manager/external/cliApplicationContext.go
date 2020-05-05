package external

import (
	"github.com/halprin/email-conceal/manager/controller"
	"github.com/halprin/email-conceal/manager/external/lib"
	"github.com/halprin/email-conceal/manager/gateways"
	"github.com/halprin/email-conceal/manager/usecases"
	"os"
)

type CliApplicationContext struct{}

func (cliAppContext *CliApplicationContext) ConcealEmailController(cliArguments []string) string {
	return controller.CliConcealEmailController(cliArguments, cliAppContext)
}

func (cliAppContext *CliApplicationContext) EnvironmentGateway(key string) string {
	return gateways.OsEnvEnvironmentGateway(key, cliAppContext)
}

func (cliAppContext *CliApplicationContext) AddConcealEmailUsecase(email string) (string, error) {
	return usecases.AddConcealEmailUsecase(email, cliAppContext)
}

func (cliAppContext *CliApplicationContext) GenerateRandomUuid() string {
	return lib.GenerateGoogleRandomUuid(cliAppContext)
}

func (cliAppContext *CliApplicationContext) Exit(returnCode int) {
	os.Exit(returnCode)
}
