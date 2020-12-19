package sqsQueue

import (
	"github.com/halprin/email-conceal/forwarder/controllers"
	"github.com/halprin/email-conceal/forwarder/gateways"
	"github.com/halprin/email-conceal/forwarder/usecases"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/controllers/forwardEmail"
	"github.com/halprin/email-conceal/src/external/lib"
	"github.com/halprin/email-conceal/src/gateways/dynamodb"
	"github.com/halprin/email-conceal/src/gateways/osEnvironmentVariable"
	concealEmailUsecase "github.com/halprin/email-conceal/src/usecases/concealEmail"
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


func init() {
	var applicationContext = context.ApplicationContext{}

	//controllers
	applicationContext.Bind(func() forwardEmail.ForwardEmail {
		return forwardEmail.SqsQueueForwardController{}
	})

	//usecases
	applicationContext.Bind(func() concealEmailUsecase.ConcealEmailUsecase {
		return concealEmailUsecase.ConcealEmailUsecaseImpl{}
	})

	//gateways
	applicationContext.Bind(func() concealEmailUsecase.ConcealEmailGateway {
		return dynamodb.DynamoDbGateway{}
	})

	applicationContext.Bind(func() context.EnvironmentGateway {
		return osEnvironmentVariable.OsEnvironmentGateway{}
	})

	//libraries
	applicationContext.Bind(func() context.UuidLibrary {
		return lib.GoogleUuid{}
	})
}