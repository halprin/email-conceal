package controllers

import (
	"github.com/halprin/email-conceal/forwarder/context"
	"github.com/halprin/email-conceal/forwarder/external/lib/errors"
	"net/http"
	"testing"
)

type TestGinContext struct {
	ReturnErrorWhenBindingJson error
	UrlToAddWhenBindingJson string
	StringMethodCalled bool
	ResponseStatusReturned int
}

func(testGinContext *TestGinContext) BindJSON(obj interface{}) error {
	if testGinContext.ReturnErrorWhenBindingJson != nil {
		return testGinContext.ReturnErrorWhenBindingJson
	}

	requestBody := obj.(*RestApiRequestBody)
	requestBody.Url = testGinContext.UrlToAddWhenBindingJson

	return nil
}

func(testGinContext *TestGinContext) String(code int, format string, values ...interface{}) {
	testGinContext.StringMethodCalled = true
	testGinContext.ResponseStatusReturned = code
}

func TestRestApiForwardEmailFailsJsonParsing(t *testing.T) {
	appContext := context.TestApplicationContext{}
	requestContext := &TestGinContext{
		ReturnErrorWhenBindingJson: errors.New("Moof go boom"),
	}

	arguments := map[string]interface{}{
		"context": requestContext,
	}

	_ = RestApiForwardEmail(arguments, &appContext)

	if requestContext.StringMethodCalled == false {
		t.Error("The String method should have been called on the gin context")
	}
	if requestContext.ResponseStatusReturned != http.StatusBadRequest {
		t.Error("The response status for the gin context should have been StatusBadRequest")
	}
	if appContext.ReceivedForwardEmailUsecaseArguments != "" {
		t.Error("The EmailUsecase was called, it shouldn't have been")
	}
}

func TestRestApiForwardEmailFailsInTheUsecase(t *testing.T) {
	appContext := context.TestApplicationContext{
		ReturnErrorForwardEmailUsecase: errors.New("Moof go boom"),
	}
	testEmailUrl := "http://dogcow.com"
	requestContext := &TestGinContext{
		UrlToAddWhenBindingJson: testEmailUrl,
	}

	arguments := map[string]interface{}{
		"context": requestContext,
	}

	_ = RestApiForwardEmail(arguments, &appContext)

	if requestContext.StringMethodCalled == false {
		t.Error("The String method should have been called on the gin context")
	}
	if requestContext.ResponseStatusReturned != http.StatusInternalServerError {
		t.Error("The response status for the gin context should have been StatusInternalServerError")
	}
	if appContext.ReceivedForwardEmailUsecaseArguments != testEmailUrl {
		t.Error("The EmailUsecase should have been called with " + testEmailUrl + ", it wasn't")
	}
}

func TestRestApiForwardEmailWorks(t *testing.T) {
	appContext := context.TestApplicationContext{}
	testEmailUrl := "http://dogcow.com"
	requestContext := &TestGinContext{
		UrlToAddWhenBindingJson: testEmailUrl,
	}

	arguments := map[string]interface{}{
		"context": requestContext,
	}

	_ = RestApiForwardEmail(arguments, &appContext)

	if requestContext.StringMethodCalled == false {
		t.Error("The String method should have been called on the gin context")
	}
	if requestContext.ResponseStatusReturned != http.StatusCreated {
		t.Error("The response status for the gin context should have been StatusCreated")
	}
	if appContext.ReceivedForwardEmailUsecaseArguments != testEmailUrl {
		t.Error("The EmailUsecase should have been called with " + testEmailUrl + ", it wasn't")
	}
}
