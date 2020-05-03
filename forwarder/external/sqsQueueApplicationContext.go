package external

import (
	"github.com/halprin/email-conceal/forwarder/controllers"
	"github.com/halprin/email-conceal/forwarder/gateways"
	"github.com/halprin/email-conceal/forwarder/usecases"
	"os"
)

type SqsQueueApplicationContext struct{}

func (appContext *SqsQueueApplicationContext) ForwardEmailController(arguments map[string]interface{}) error {
	return controllers.SqsQueueForwardEmail(arguments, appContext)
}

func (appContext *SqsQueueApplicationContext) ReadEmailGateway(url string) ([]byte, error) {
	return gateways.S3ReadEmailGateway(url, appContext)
}

func (appContext *SqsQueueApplicationContext) SendEmailGateway(email []byte, recipients []string) error {
	return gateways.AwsSesSendEmailGateway(email, recipients, appContext)
}

func (appContext *SqsQueueApplicationContext) EnvironmentGateway(key string) string {
	return gateways.OsEnvEnvironmentGateway(key, appContext)
}

func (appContext *SqsQueueApplicationContext) GetRealEmailForConcealPrefix(concealPrefix string) (string, error) {
	return gateways.GetRealEmailForConcealPrefix(concealPrefix, appContext)
}

func (appContext *SqsQueueApplicationContext) ForwardEmailUsecase(url string) error {
	return usecases.ForwardEmailUsecase(url, appContext)
}

func (appContext *SqsQueueApplicationContext) Exit(returnCode int) {
	os.Exit(returnCode)
}
