package external

import (
	"github.com/halprin/email-conceal/manager/controller"
	"github.com/halprin/email-conceal/manager/external/lib"
	"github.com/halprin/email-conceal/manager/gateways"
	"github.com/halprin/email-conceal/manager/usecases"
	"os"
)

type CliApplicationContext struct{}

func (appContext *CliApplicationContext) ConcealEmailController(cliArguments []string) string {
	return controller.CliConcealEmailController(cliArguments, appContext)
}

func (appContext *CliApplicationContext) EnvironmentGateway(key string) string {
	return gateways.OsEnvEnvironmentGateway(key, appContext)
}

func (appContext *CliApplicationContext) AddConcealedEmailToActualEmailMappingGateway(concealPrefix string, actualEmail string) error {
	return gateways.AddConcealedEmailToActualEmailMapping(concealPrefix, actualEmail, appContext)
}

func (appContext *CliApplicationContext) DeleteConcealedEmailToActualEmailMappingGateway(concealPrefix string) error {
	return gateways.DeleteConcealedEmailToActualEmailMapping(concealPrefix, appContext)
}

func (appContext *CliApplicationContext) AddConcealEmailUsecase(email string) (string, error) {
	return usecases.AddConcealEmailUsecase(email, appContext)
}

func (appContext *CliApplicationContext) GenerateRandomUuid() string {
	return lib.GenerateGoogleRandomUuid(appContext)
}

func (appContext *CliApplicationContext) Exit(returnCode int) {
	os.Exit(returnCode)
}
