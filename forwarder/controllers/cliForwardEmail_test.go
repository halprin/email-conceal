package controllers

import (
	"github.com/halprin/email-conceal/forwarder/context"
	"github.com/halprin/email-conceal/forwarder/external/lib/errors"
	"testing"
)

func TestCliForwardEmailFails(t *testing.T) {
	appContext := context.TestApplicationContext{
		ReturnErrorForwardEmailUsecase: errors.New("forward e-mail usecase failed"),
	}
	arguments := map[string]interface{}{
		"url": "/path/to/email.dms",
	}

	err := CliForwardEmail(arguments, &appContext)

	if err == nil {
		t.Error("An error should have been returned from CliForwardEmail controllers")
	}

	if appContext.ReceivedExitReturnCode != 1 {
		t.Errorf("Exit should have been called with return code 1, but it wasn't called or was called with %d", appContext.ReceivedExitReturnCode)
	}
}

func TestCliForwardEmailSuccess(t *testing.T) {
	appContext := context.TestApplicationContext{}
	emailPath := "/path/to/email.dms"
	arguments := map[string]interface{}{
		"url": emailPath,
	}

	err := CliForwardEmail(arguments, &appContext)

	if err != nil {
		t.Errorf("An error was returned from CliForwardEmail when it shouldn't have been.  Error: %+v", err)
	}

	receivedForwardEmailUsecase := appContext.ReceivedForwardEmailUsecaseArguments

	if receivedForwardEmailUsecase != emailPath {
		t.Errorf("ForwardEmailUsecase's argument should have been %s, but instead it was %s", emailPath, receivedForwardEmailUsecase)
	}
}
