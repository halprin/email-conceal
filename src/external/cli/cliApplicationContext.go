package cli

import (
	"github.com/halprin/email-conceal/forwarder/controllers"
	"github.com/halprin/email-conceal/forwarder/gateways"
	"github.com/halprin/email-conceal/forwarder/usecases"
	"github.com/halprin/email-conceal/src/context"
	concealEmailController "github.com/halprin/email-conceal/src/controllers/concealEmail"
	"github.com/halprin/email-conceal/src/controllers/forwardEmail"
	"github.com/halprin/email-conceal/src/external/lib"
	"github.com/halprin/email-conceal/src/gateways/dynamodb"
	"github.com/halprin/email-conceal/src/gateways/osEnvironmentVariable"
	concealEmailUsecase "github.com/halprin/email-conceal/src/usecases/concealEmail"
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
	applicationContext.Bind(func() forwardEmail.ForwardEmail {
		return forwardEmail.CliForwardController{}
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
