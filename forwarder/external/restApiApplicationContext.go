package external

import (
	"github.com/halprin/email-conceal/forwarder/controllers"
	"github.com/halprin/email-conceal/forwarder/gateways"
	"github.com/halprin/email-conceal/forwarder/usecases"
	"os"
)

type RestApiApplicationContext struct{}

func (appContext *RestApiApplicationContext) ForwardEmailController(arguments map[string]interface{}) error {
	return controllers.RestApiForwardEmail(arguments, appContext)
}

func (appContext *RestApiApplicationContext) ReadEmailGateway(url string) ([]byte, error) {
	return gateways.FileReadEmailGateway(url, appContext)
}

func (appContext *RestApiApplicationContext) SendEmailGateway(email []byte) error {
	return gateways.AwsSesSendEmailGateway(email, appContext)
}

func (appContext *RestApiApplicationContext) EnvironmentGateway(key string) string {
	return gateways.OsEnvEnvironmentGateway(key, appContext)
}

func (appContext *RestApiApplicationContext) ForwardEmailUsecase(url string) error {
	return usecases.ForwardEmailUsecase(url, appContext)
}

func (appContext *RestApiApplicationContext) Exit(returnCode int) {
	os.Exit(returnCode)
}
