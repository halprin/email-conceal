package external

import (
	"github.com/halprin/email-conceal/forwarder/controllers"
	"github.com/halprin/email-conceal/forwarder/gateways"
	"github.com/halprin/email-conceal/forwarder/usecases"
	"os"
)

type CliApplicationContext struct{}

func (cliAppContext *CliApplicationContext) ForwardEmailController(arguments map[string]interface{}) error {
	return controllers.CliForwardEmail(arguments, cliAppContext)
}

func (cliAppContext *CliApplicationContext) ReadEmailGateway(url string) ([]byte, error) {
	return gateways.FileReadEmailGateway(url, cliAppContext)
}

func (cliAppContext *CliApplicationContext) SendEmailGateway(email []byte) error {
	return gateways.AwsSesSendEmailGateway(email, cliAppContext)
}

func (cliAppContext *CliApplicationContext) EnvironmentGateway(key string) string {
	return gateways.OsEnvEnvironmentGateway(key, cliAppContext)
}

func (cliAppContext *CliApplicationContext) ForwardEmailUsecase(url string) error {
	return usecases.ForwardEmailUsecase(url, cliAppContext)
}

func (cliAppContext *CliApplicationContext) Exit(returnCode int) {
	os.Exit(returnCode)
}
