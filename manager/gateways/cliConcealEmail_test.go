package gateways

import (
	"github.com/halprin/email-conceal/context"
	"github.com/halprin/email-conceal/entities"
	"testing"
)

func TestCliConcealEmailGateway(t *testing.T) {
	testApplicationContext := &context.TestApplicationContext{
		ReturnFromConcealEmailUsecase: "concealed@asdf.com",
	}

	sourceEmail := "dogcow@apple.com"
	var arguments = []string{"program_invoked", sourceEmail}
	actualConcealedEmail := CliConcealEmailGateway(arguments, testApplicationContext)

	if testApplicationContext.ReceivedConcealEmailUsecaseEmail != sourceEmail {
		t.Errorf("The parsed source e-mail %s was not the passed in e-mail %s", testApplicationContext.ReceivedConcealEmailUsecaseEmail, sourceEmail)
	}

	if actualConcealedEmail != testApplicationContext.ReturnFromConcealEmailUsecase {
		t.Errorf("The concealed e-mail %s generated wasn't passed back completely, instead %s was returned", testApplicationContext.ReturnFromConcealEmailUsecase, actualConcealedEmail)
	}
}

func TestCliConcealEmailGatewayNegative(t *testing.T) {
	testApplicationContext := &context.TestApplicationContext{
		ReturnErrorFromConcealEmailUsecase: entities.InvalidEmailAddressError,
	}

	var arguments = []string{"program_invoked", "dogcow@apple.com"}
	CliConcealEmailGateway(arguments, testApplicationContext)

	const expectedReturnCode = 1
	if testApplicationContext.ReceivedExitReturnCode != expectedReturnCode {
		t.Errorf("The program should have decided to exit with a return code of %d, but instead %d was returned", expectedReturnCode, testApplicationContext.ReceivedExitReturnCode)
	}
}
