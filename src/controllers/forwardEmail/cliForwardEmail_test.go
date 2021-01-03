package forwardEmail

import (
	"fmt"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"github.com/halprin/email-conceal/src/usecases/forwardEmail"
	"testing"
)

var cliController = CliForwardController{}

func TestCliForwardEmailFails(t *testing.T) {

	testForwardEmailUsecase := TestForwardEmailUsecase{
		ForwardEmailReturnError: errors.New("forward e-mail usecase failed"),
	}
	testAppContext.Bind(func() forwardEmail.ForwardEmailUsecase {
		return &testForwardEmailUsecase
	})

	arguments := map[string]interface{}{
		"url": "/path/to/email.dms",
	}

	fmt.Println("About to go!")
	err := cliController.ForwardEmail(arguments)
	fmt.Println("Let's do this!")

	if err == nil {
		t.Error("An error should have been returned from CliForwardEmail controller")
	}
}

func TestCliForwardEmailSuccess(t *testing.T) {
	emailPath := "/path/to/email.dms"
	arguments := map[string]interface{}{
		"url": emailPath,
	}

	testForwardEmailUsecase := TestForwardEmailUsecase{}
	testAppContext.Bind(func() forwardEmail.ForwardEmailUsecase {
		return &testForwardEmailUsecase
	})

	err := cliController.ForwardEmail(arguments)

	if err != nil {
		t.Errorf("An error was returned from CliForwardEmail when it shouldn't have been.  Error: %+v", err)
	}

	receivedForwardEmailUsecase := testForwardEmailUsecase.ForwardEmailUri

	if receivedForwardEmailUsecase != emailPath {
		t.Errorf("ForwardEmailUsecase's argument should have been %s, but instead it was %s", emailPath, receivedForwardEmailUsecase)
	}
}
