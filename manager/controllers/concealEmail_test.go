package controllers

import (
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/entities"
	"github.com/halprin/email-conceal/manager/external/lib/errors"
	"github.com/halprin/email-conceal/manager/usecases"
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

func TestConcealEmailControllerSuccess(t *testing.T) {
	concealedEmail := "concealed@asdf.com"

	testUsecase := TestConcealEmailUsecase{
		AddReturnConcealEmail: concealedEmail,
	}
	testAppContext.Bind(func() usecases.ConcealEmailUsecase {
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
	testAppContext.Bind(func() usecases.ConcealEmailUsecase {
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
	testAppContext.Bind(func() usecases.ConcealEmailUsecase {
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
	testAppContext.Bind(func() usecases.ConcealEmailUsecase {
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

	testAppContext.Bind(func() usecases.ConcealEmailUsecase {
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

	testAppContext.Bind(func() usecases.ConcealEmailUsecase {
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

	testAppContext.Bind(func() usecases.ConcealEmailUsecase {
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
