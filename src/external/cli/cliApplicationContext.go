package cli

import (
	"github.com/halprin/email-conceal/src/context"
	forwardEmailController "github.com/halprin/email-conceal/src/controllers/forwardEmail"
	"github.com/halprin/email-conceal/src/external/lib"
	"github.com/halprin/email-conceal/src/gateways/dynamodb"
	"github.com/halprin/email-conceal/src/gateways/osEnvironmentVariable"
	forwardEmailUsecase "github.com/halprin/email-conceal/src/usecases/forwardEmail"
	"os"
)

type CliApplicationContext struct{}

func (cliAppContext *CliApplicationContext) ForwardEmailController(arguments map[string]interface{}) error {
	return controllers.CliForwardEmail(arguments, cliAppContext)
}

func (cliAppContext *CliApplicationContext) ReadEmailGateway(url string) ([]byte, error) {
	return gateways.FileReadEmailGateway(url, cliAppContext)
}

func (cliAppContext *CliApplicationContext) SendEmailGateway(email []byte, recipients []string) error {
	return gateways.AwsSesSendEmailGateway(email, recipients, cliAppContext)
}

func (cliAppContext *CliApplicationContext) EnvironmentGateway(key string) string {
	return gateways.OsEnvEnvironmentGateway(key, cliAppContext)
}

func (cliAppContext *CliApplicationContext) GetRealEmailForConcealPrefix(concealPrefix string) (string, error) {
	return gateways.GetRealEmailForConcealPrefix(concealPrefix, cliAppContext)
}

func (cliAppContext *CliApplicationContext) ForwardEmailUsecase(url string) error {
	return usecases.ForwardEmailUsecase(url, cliAppContext)
}

func (cliAppContext *CliApplicationContext) Exit(returnCode int) {
	os.Exit(returnCode)
}

func init() {
	var applicationContext = context.ApplicationContext{}

	//controllers
	applicationContext.Bind(func() forwardEmailController.ForwardEmail {
		return forwardEmailController.CliForwardController{}
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
