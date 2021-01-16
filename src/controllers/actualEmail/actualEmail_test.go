package actualEmail

import (
	"github.com/halprin/email-conceal/src/context"
	actualEmailUsecase2 "github.com/halprin/email-conceal/src/usecases/actualEmail"
	"net/http"
	"testing"
)


var controller = ActualEmailController{}
var testAppContext = context.ApplicationContext{}

type TestActualEmailUsecase struct {
	AddReceiveActualEmail string
	AddReturnError        error
}

func (testUsecase *TestActualEmailUsecase) Add(actualEmail string) error {
	testUsecase.AddReceiveActualEmail = actualEmail
	return testUsecase.AddReturnError
}

func TestActualEmailControllerSuccess(t *testing.T) {
	testAppContext.Reset()
	testUsecase := TestActualEmailUsecase{}
	testAppContext.Bind(func() actualEmailUsecase2.ActualEmailUsecase {
		return &testUsecase
	})
	controller.Init()

	status, body := controller.Add(map[string]interface{}{
		"email": "moof@dogcow.com",
	})

	if status != http.StatusCreated {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusCreated)
	}

	if body != nil {
		t.Error("The body was not nil, but it should have been")
	}
}
