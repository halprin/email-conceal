package concealEmail

import (
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/entities"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"github.com/halprin/email-conceal/src/usecases"
	"github.com/halprin/email-conceal/src/usecases/concealEmail"
	"net/http"
	"testing"
)


var controller = ConcealEmailController{}
var testAppContext = context.ApplicationContext{}

type TestConcealEmailUsecase struct {
	AddReceiveSourceEmail string
	AddReceiveDescription *string
	AddReturnConcealEmail string
	AddReturnError        error

	DeleteReceiveConcealEmailPrefix string
	DeleteReturnError               error

	AddDescriptionReceiveConcealEmailPrefix string
	AddDescriptionReceiveDescription        string
	AddDescriptionReturnError               error

	DeleteDescriptionReceiveConcealEmailPrefix string
	DeleteDescriptionReturnError               error
}

func (testUsecase *TestConcealEmailUsecase) Add(sourceEmail string, description *string) (string, error) {
	testUsecase.AddReceiveSourceEmail = sourceEmail
	testUsecase.AddReceiveDescription = description
	return testUsecase.AddReturnConcealEmail, testUsecase.AddReturnError
}

func (testUsecase *TestConcealEmailUsecase) Delete(concealedEmailPrefix string) error {
	testUsecase.DeleteReceiveConcealEmailPrefix = concealedEmailPrefix
	return testUsecase.DeleteReturnError
}

func (testUsecase *TestConcealEmailUsecase) AddDescriptionToExistingEmail(concealedEmailPrefix string, description string) error {
	testUsecase.AddDescriptionReceiveConcealEmailPrefix = concealedEmailPrefix
	testUsecase.AddDescriptionReceiveDescription = description
	return testUsecase.AddDescriptionReturnError
}

func (testUsecase *TestConcealEmailUsecase) DeleteDescriptionFromExistingEmail(concealedEmailPrefix string) error {
	testUsecase.DeleteDescriptionReceiveConcealEmailPrefix = concealedEmailPrefix
	return testUsecase.DeleteDescriptionReturnError
}

func TestConcealEmailControllerSuccess(t *testing.T) {
	concealedEmail := "concealed@asdf.com"

	testUsecase := TestConcealEmailUsecase{
		AddReturnConcealEmail: concealedEmail,
	}
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &testUsecase
	})

	sourceEmail := "dogcow@apple.com"
	var arguments = map[string]interface{}{
		"email": sourceEmail,
	}

	status, body := controller.Add(arguments)
	actualConcealedEmail := body["concealedEmail"]

	if testUsecase.AddReceiveSourceEmail != sourceEmail {
		t.Errorf("The parsed source e-mail %s was not the passed in e-mail %s", testUsecase.AddReceiveSourceEmail, sourceEmail)
	}

	if actualConcealedEmail != testUsecase.AddReturnConcealEmail {
		t.Errorf("The concealed e-mail %s generated wasn't passed back completely, instead %s was returned", testUsecase.AddReturnConcealEmail, actualConcealedEmail)
	}

	if status != http.StatusCreated {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusCreated)
	}
}

func TestConcealEmailControllerBadEmailType(t *testing.T) {

	testUsecase := TestConcealEmailUsecase{}
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &testUsecase
	})

	var arguments = map[string]interface{}{
		"email": 3,
	}

	status, _ := controller.Add(arguments)

	if testUsecase.AddReceiveSourceEmail != "" {
		t.Errorf("The usecase was called, but it shouldn't have been")
	}

	if status != http.StatusBadRequest {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusBadRequest)
	}
}

func TestConcealEmailControllerInvalidEmail(t *testing.T) {

	testUsecase := TestConcealEmailUsecase{
		AddReturnError: entities.InvalidEmailAddressError,
	}
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &testUsecase
	})

	sourceEmail := "dogcow"
	var arguments = map[string]interface{}{
		"email": sourceEmail,
	}

	status, _ := controller.Add(arguments)

	if testUsecase.AddReceiveSourceEmail != sourceEmail {
		t.Errorf("The parsed source e-mail %s was not the passed in e-mail %s", testUsecase.AddReceiveSourceEmail, sourceEmail)
	}

	if status != http.StatusBadRequest {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusBadRequest)
	}
}

func TestConcealEmailControllerUnknownError(t *testing.T) {

	testUsecase := TestConcealEmailUsecase{
		AddReturnError: errors.New("some other error"),
	}
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &testUsecase
	})

	sourceEmail := "dogcow@apple.com"
	var arguments = map[string]interface{}{
		"email": sourceEmail,
	}

	status, _ := controller.Add(arguments)

	if testUsecase.AddReceiveSourceEmail != sourceEmail {
		t.Errorf("The parsed source e-mail %s was not the passed in e-mail %s", testUsecase.AddReceiveSourceEmail, sourceEmail)
	}

	if status != http.StatusInternalServerError {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusInternalServerError)
	}
}

func TestDeleteConcealEmailControllerSuccess(t *testing.T) {

	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{}
	})

	var arguments = map[string]interface{}{
		"concealEmailId": "dogcow",
	}

	status, body := controller.Delete(arguments)

	if status != http.StatusNoContent {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusNoContent)
	}

	if len(body) != 0 {
		t.Errorf("The returned status response body wasn't empty; it should've been")
	}
}

func TestDeleteConcealEmailControllerBadInput(t *testing.T) {

	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{}
	})

	var arguments = map[string]interface{}{
		"concealEmailId": 3,
	}

	status, body := controller.Delete(arguments)

	if status != http.StatusBadRequest {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusBadRequest)
	}

	_, exists := body["error"]
	if !exists {
		t.Errorf("An error is missing from the response body; it should've been there")
	}
}

func TestDeleteConcealEmailControllerFailedDelete(t *testing.T) {

	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{
			DeleteReturnError: errors.New("moof! go boom"),
		}
	})

	var arguments = map[string]interface{}{
		"concealEmailId": "dogcow",
	}

	status, body := controller.Delete(arguments)

	if status != http.StatusInternalServerError {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusInternalServerError)
	}

	_, exists := body["error"]
	if !exists {
		t.Errorf("An error is missing from the response body; it should've been there")
	}
}

func TestUpdateConcealEmailWithNewDescription(t *testing.T) {
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{}
	})

	var arguments = map[string]interface{}{
		"concealEmailId": "an ID",
		"description": "a new description",
	}

	status, body := controller.Update(arguments)

	if status != http.StatusOK {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusOK)
	}

	bodySize := len(body)
	if bodySize != 0 {
		t.Errorf("There was data in the response body when there shouldn't be anything")
	}
}

func TestTooLongDescriptionUpdate(t *testing.T) {
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{
			AddDescriptionReturnError: entities.DescriptionTooLongError,
		}
	})

	var arguments = map[string]interface{}{
		"concealEmailId": "an ID",
		"description": "this value doesn't matter",
	}

	status, body := controller.Update(arguments)

	if status != http.StatusBadRequest {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusBadRequest)
	}

	_, exists := body["error"]
	if !exists {
		t.Errorf("An error is missing from the response body; it should've been there")
	}
}

func TestDescriptionUpdateFailedForUnkownReason(t *testing.T) {
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{
			AddDescriptionReturnError: errors.New("Unknown error"),
		}
	})

	var arguments = map[string]interface{}{
		"concealEmailId": "an ID",
		"description": "this value doesn't matter",
	}

	status, body := controller.Update(arguments)

	if status != http.StatusInternalServerError {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusInternalServerError)
	}

	_, exists := body["error"]
	if !exists {
		t.Errorf("An error is missing from the response body; it should've been there")
	}
}

func TestDescriptionUpdateWithDelete(t *testing.T) {

	concealEmailUsecase := TestConcealEmailUsecase{}
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &concealEmailUsecase
	})

	conceaEmailId := "an ID"
	var arguments = map[string]interface{}{
		"concealEmailId": conceaEmailId,
		"description": "",  //empty on purpose
	}

	status, body := controller.Update(arguments)

	if concealEmailUsecase.AddDescriptionReceiveConcealEmailPrefix != "" && concealEmailUsecase.DeleteDescriptionReceiveConcealEmailPrefix != conceaEmailId {
		t.Errorf("The wrong usecase was called.  The delete usecase should have been called.")
	}

	if status != http.StatusOK {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusOK)
	}

	_, exists := body["error"]
	if exists {
		t.Errorf("An error was returned in the response body; it shouldn't be there")
	}
}

func TestUpdateFailedWithConcealEmailNotExist(t *testing.T) {

	conceaEmailId := "an ID"
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{
			DeleteDescriptionReturnError: usecases.ConcealEmailNotExistError{
				ConcealEmailId: conceaEmailId,
			},
		}
	})

	var arguments = map[string]interface{}{
		"concealEmailId": conceaEmailId,
		"description": "",  //empty on purpose
	}

	status, body := controller.Update(arguments)

	if status != http.StatusNotFound {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusNotFound)
	}

	_, exists := body["error"]
	if !exists {
		t.Errorf("An error is missing from the response body; it should've been there")
	}
}
