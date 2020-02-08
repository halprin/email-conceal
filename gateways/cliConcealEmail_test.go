package gateways

import (
	"github.com/halprin/email-conceal/entities"
	"testing"
)

var concealedEmail string
var concealEmailError error = nil
var receivedSourceEmail string
var receivedReturnCode int

func TestCliConcealEmailGateway(t *testing.T) {
	concealedEmail = "concealed@asdf.com"
	concealEmailError = nil
	sourceEmail := "dogcow@apple.com"

	var arguments = []string {"program_invoked", sourceEmail}
	actualConcealedEmail := CliConcealEmailGateway(arguments, testApplicationContext{})

	if receivedSourceEmail != sourceEmail {
		t.Errorf("The parsed source e-mail %s was not the passed in e-mail %s", receivedSourceEmail, sourceEmail)
	}

	if actualConcealedEmail != concealedEmail {
		t.Errorf("The concealed e-mail %s generated wasn't passed back completely, instead %s was returned", concealedEmail, actualConcealedEmail)
	}
}

func TestCliConcealEmailGatewayNegative(t *testing.T) {
	concealedEmail = ""
	concealEmailError = entities.InvalidEmailAddressError

	var arguments = []string {"program_invoked", "dogcow@apple.com"}
	CliConcealEmailGateway(arguments, testApplicationContext{})

	const expectedReturnCode = 1
	if receivedReturnCode != expectedReturnCode {
		t.Errorf("The program should have decided to exit with a return code of %d, but instead %d was returned", expectedReturnCode, receivedReturnCode)
	}
}

//Test application context
type testApplicationContext struct{}

func (appContext testApplicationContext) ConcealEmailGateway(cliArguments []string) string {
	return ""
}

func (appContext testApplicationContext) ConcealEmailUsecase(email string) (string, error) {
	receivedSourceEmail = email
	return concealedEmail, concealEmailError
}

func (appContext testApplicationContext) GenerateRandomUuid() string {
	return ""
}

func (appContext testApplicationContext) Exit(returnCode int) {
	receivedReturnCode = returnCode
}
