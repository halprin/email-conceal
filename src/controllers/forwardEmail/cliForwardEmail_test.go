package forwardEmail

import (
	"fmt"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"github.com/halprin/email-conceal/src/usecases/forwardEmail"
	"testing"
)

var controller = CliForwardController{}
var testAppContext = context.ApplicationContext{}

type TestForwardEmailUsecase struct {
	ForwardEmailUri         string
	ForwardEmailReturnError error
}

func (testUsecase *TestForwardEmailUsecase) ForwardEmail(url string) error {
	testUsecase.ForwardEmailUri = url

	return testUsecase.ForwardEmailReturnError
}

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
	err := controller.ForwardEmail(arguments)
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

	err := controller.ForwardEmail(arguments)

	if err != nil {
		t.Errorf("An error was returned from CliForwardEmail when it shouldn't have been.  Error: %+v", err)
	}

	receivedForwardEmailUsecase := testForwardEmailUsecase.ForwardEmailUri

	if receivedForwardEmailUsecase != emailPath {
		t.Errorf("ForwardEmailUsecase's argument should have been %s, but instead it was %s", emailPath, receivedForwardEmailUsecase)
	}
}
