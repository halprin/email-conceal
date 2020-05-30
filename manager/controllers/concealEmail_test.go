package controllers

import (
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/entities"
	"github.com/halprin/email-conceal/manager/external/lib/errors"
	"net/http"
	"testing"
)

func TestConcealEmailGateway(t *testing.T) {
	concealedEmail := "concealed@asdf.com"
	testApplicationContext := &context.TestApplicationContext{
		ReturnFromConcealEmailUsecase: concealedEmail,
	}

	sourceEmail := "dogcow@apple.com"
	var arguments = map[string]interface{}{
		"email": sourceEmail,
	}

	status, body := HttpConcealEmailController(arguments, testApplicationContext)
	actualConcealedEmail := body["concealedEmail"]

	if testApplicationContext.ReceivedConcealEmailUsecaseEmail != sourceEmail {
		t.Errorf("The parsed source e-mail %s was not the passed in e-mail %s", testApplicationContext.ReceivedConcealEmailUsecaseEmail, sourceEmail)
	}

	if actualConcealedEmail != testApplicationContext.ReturnFromConcealEmailUsecase {
		t.Errorf("The concealed e-mail %s generated wasn't passed back completely, instead %s was returned", testApplicationContext.ReturnFromConcealEmailUsecase, actualConcealedEmail)
	}

	if status != http.StatusCreated {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusCreated)
	}
}

func TestConcealEmailGatewayBadEmailType(t *testing.T) {
	testApplicationContext := &context.TestApplicationContext{}

	var arguments = map[string]interface{}{
		"email": 3,
	}

	status, _ := HttpConcealEmailController(arguments, testApplicationContext)

	if testApplicationContext.ReceivedConcealEmailUsecaseEmail != "" {
		t.Errorf("The usecase was called, but it shouldn't have been")
	}

	if status != http.StatusBadRequest {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusBadRequest)
	}
}

func TestConcealEmailGatewayInvalidEmail(t *testing.T) {
	testApplicationContext := &context.TestApplicationContext{
		ReturnErrorFromConcealEmailUsecase: entities.InvalidEmailAddressError,
	}

	sourceEmail := "dogcow"
	var arguments = map[string]interface{}{
		"email": sourceEmail,
	}

	status, _ := HttpConcealEmailController(arguments, testApplicationContext)

	if testApplicationContext.ReceivedConcealEmailUsecaseEmail != sourceEmail {
		t.Errorf("The parsed source e-mail %s was not the passed in e-mail %s", testApplicationContext.ReceivedConcealEmailUsecaseEmail, sourceEmail)
	}

	if status != http.StatusBadRequest {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusBadRequest)
	}
}

func TestConcealEmailGatewayUnknownError(t *testing.T) {
	testApplicationContext := &context.TestApplicationContext{
		ReturnErrorFromConcealEmailUsecase: errors.New("some other error"),
	}

	sourceEmail := "dogcow@apple.com"
	var arguments = map[string]interface{}{
		"email": sourceEmail,
	}

	status, _ := HttpConcealEmailController(arguments, testApplicationContext)

	if testApplicationContext.ReceivedConcealEmailUsecaseEmail != sourceEmail {
		t.Errorf("The parsed source e-mail %s was not the passed in e-mail %s", testApplicationContext.ReceivedConcealEmailUsecaseEmail, sourceEmail)
	}

	if status != http.StatusInternalServerError {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusInternalServerError)
	}
}

func TestDeleteConcealEmailGateway(t *testing.T) {
	testApplicationContext := &context.TestApplicationContext{}

	var arguments = map[string]interface{}{
		"concealEmailId": "dogcow",
	}

	status, body := HttpDeleteConcealEmailController(arguments, testApplicationContext)

	if status != http.StatusNoContent {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusNoContent)
	}

	if len(body) != 0 {
		t.Errorf("The returned status response body wasn't empty; it should've been")
	}
}

func TestDeleteConcealEmailGatewayBadInput(t *testing.T) {
	testApplicationContext := &context.TestApplicationContext{}

	var arguments = map[string]interface{}{
		"concealEmailId": 3,
	}

	status, body := HttpDeleteConcealEmailController(arguments, testApplicationContext)

	if status != http.StatusBadRequest {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusBadRequest)
	}

	_, exists := body["error"]
	if !exists {
		t.Errorf("An error is missing from the response body; it should've been there")
	}
}

func TestDeleteConcealEmailGatewayFailedDelete(t *testing.T) {
	testApplicationContext := &context.TestApplicationContext{
		ReturnErrorFromDeleteConcealEmailUsecase: errors.New("moof! go boom"),
	}

	var arguments = map[string]interface{}{
		"concealEmailId": "dogcow",
	}

	status, body := HttpDeleteConcealEmailController(arguments, testApplicationContext)

	if status != http.StatusInternalServerError {
		t.Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusInternalServerError)
	}

	_, exists := body["error"]
	if !exists {
		t.Errorf("An error is missing from the response body; it should've been there")
	}
}
