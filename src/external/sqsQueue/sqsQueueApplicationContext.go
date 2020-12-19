package sqsQueue

import (
	"github.com/halprin/email-conceal/src/context"
	forwardEmailController "github.com/halprin/email-conceal/src/controllers/forwardEmail"
	"github.com/halprin/email-conceal/src/external/lib"
	"github.com/halprin/email-conceal/src/gateways/dynamodb"
	"github.com/halprin/email-conceal/src/gateways/osEnvironmentVariable"
	forwardEmailUsecase "github.com/halprin/email-conceal/src/usecases/forwardEmail"
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
	applicationContext.Bind(func() forwardEmailController.ForwardEmail {
		return forwardEmailController.SqsQueueForwardController{}
	})

	//usecases
	applicationContext.Bind(func() forwardEmailUsecase.ForwardEmailUsecase {
		return forwardEmailUsecase.ForwardEmailUsecaseImpl{}
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