package usecases

import (
	"github.com/halprin/email-conceal/forwarder/context"
	"github.com/halprin/email-conceal/forwarder/external/lib/errors"
	"testing"
)

func TestForwardEmailUsecaseWithFailingToReadEmail(t *testing.T) {
	errorFromGateway := errors.New("something bad happened")
	appContext := context.TestApplicationContext{
		ReturnFromReadEmailGateway:           nil,
		ReturnErrorFromReadEmailGateway:      errorFromGateway,
	}

	testUrl := "https://email.com"
	err := ForwardEmailUsecase(testUrl, &appContext)

	if !errors.Is(err, NewUnableToReadEmailError(testUrl, errorFromGateway)) {
		t.Errorf("An UnableToReadEmailError should have been returned from ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}
}
