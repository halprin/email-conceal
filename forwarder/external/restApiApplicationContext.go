package external

import (
	"github.com/halprin/email-conceal/forwarder/gateways"
	"github.com/halprin/email-conceal/forwarder/usecases"
	"os"
)

type RestApiApplicationContext struct{}

func (cliAppContext *RestApiApplicationContext) ForwardEmailGateway(arguments map[string]interface{}) error {
	return gateways.RestApiForwardEmail(arguments, cliAppContext)
}

func (cliAppContext *RestApiApplicationContext) ReadEmailGateway(url string) ([]byte, error) {
	return gateways.FileReadEmailGateway(url, cliAppContext)
}

func (cliAppContext *RestApiApplicationContext) SendEmailGateway(email []byte) error {
	return gateways.AwsSesSendEmailGateway(email, cliAppContext)
}

func (cliAppContext *RestApiApplicationContext) EnvironmentGateway(key string) string {
	return gateways.OsEnvEnvironmentGateway(key, cliAppContext)
}

func (cliAppContext *RestApiApplicationContext) ForwardEmailUsecase(url string) error {
	return usecases.ForwardEmailUsecase(url, cliAppContext)
}

func (cliAppContext *RestApiApplicationContext) Exit(returnCode int) {
	os.Exit(returnCode)
}
