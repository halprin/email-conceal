package concealEmail

import (
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/entities"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"github.com/halprin/email-conceal/src/usecases"
	"github.com/halprin/email-conceal/src/usecases/concealEmail"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

var controller = ConcealEmailController{}
var testAppContext = context.ApplicationContext{}

type ConcealEmailControllerTestSuite struct {
	suite.Suite
}

func (suite *ConcealEmailControllerTestSuite) SetupTest() {

}

func (suite *ConcealEmailControllerTestSuite) TestConcealEmailControllerSuccess() {
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
		suite.T().Errorf("The parsed source e-mail %s was not the passed in e-mail %s", testUsecase.AddReceiveSourceEmail, sourceEmail)
	}

	if actualConcealedEmail != testUsecase.AddReturnConcealEmail {
		suite.T().Errorf("The concealed e-mail %s generated wasn't passed back completely, instead %s was returned", testUsecase.AddReturnConcealEmail, actualConcealedEmail)
	}

	if status != http.StatusCreated {
		suite.T().Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusCreated)
	}
}

func (suite *ConcealEmailControllerTestSuite) TestConcealEmailControllerBadEmailType() {

	testUsecase := TestConcealEmailUsecase{}
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &testUsecase
	})

	var arguments = map[string]interface{}{
		"email": 3,
	}

	status, _ := controller.Add(arguments)

	if testUsecase.AddReceiveSourceEmail != "" {
		suite.T().Errorf("The usecase was called, but it shouldn't have been")
	}

	if status != http.StatusBadRequest {
		suite.T().Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusBadRequest)
	}
}

func (suite *ConcealEmailControllerTestSuite) TestConcealEmailControllerInvalidEmail() {

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
		suite.T().Errorf("The parsed source e-mail %s was not the passed in e-mail %s", testUsecase.AddReceiveSourceEmail, sourceEmail)
	}

	if status != http.StatusBadRequest {
		suite.T().Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusBadRequest)
	}
}

func (suite *ConcealEmailControllerTestSuite) TestConcealEmailControllerUnknownError() {

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
		suite.T().Errorf("The parsed source e-mail %s was not the passed in e-mail %s", testUsecase.AddReceiveSourceEmail, sourceEmail)
	}

	if status != http.StatusInternalServerError {
		suite.T().Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusInternalServerError)
	}
}

func (suite *ConcealEmailControllerTestSuite) TestDeleteConcealEmailControllerSuccess() {

	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{}
	})

	var arguments = map[string]interface{}{
		"concealEmailId": "dogcow",
	}

	status, body := controller.Delete(arguments)

	if status != http.StatusNoContent {
		suite.T().Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusNoContent)
	}

	if len(body) != 0 {
		suite.T().Errorf("The returned status response body wasn't empty; it should've been")
	}
}

func (suite *ConcealEmailControllerTestSuite) TestDeleteConcealEmailControllerBadInput() {

	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{}
	})

	var arguments = map[string]interface{}{
		"concealEmailId": 3,
	}

	status, body := controller.Delete(arguments)

	if status != http.StatusBadRequest {
		suite.T().Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusBadRequest)
	}

	_, exists := body["error"]
	if !exists {
		suite.T().Errorf("An error is missing from the response body; it should've been there")
	}
}

func (suite *ConcealEmailControllerTestSuite) TestDeleteConcealEmailControllerFailedDelete() {

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
		suite.T().Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusInternalServerError)
	}

	_, exists := body["error"]
	if !exists {
		suite.T().Errorf("An error is missing from the response body; it should've been there")
	}
}

func (suite *ConcealEmailControllerTestSuite) TestUpdateConcealEmailWithNewDescription() {
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{}
	})

	var arguments = map[string]interface{}{
		"concealEmailId": "an ID",
		"description":    "a new description",
	}

	status, body := controller.Update(arguments)

	if status != http.StatusOK {
		suite.T().Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusOK)
	}

	bodySize := len(body)
	if bodySize != 0 {
		suite.T().Errorf("There was data in the response body when there shouldn't be anything")
	}
}

func (suite *ConcealEmailControllerTestSuite) TestTooLongDescriptionUpdate() {
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{
			AddDescriptionReturnError: entities.DescriptionTooLongError,
		}
	})

	var arguments = map[string]interface{}{
		"concealEmailId": "an ID",
		"description":    "this value doesn't matter",
	}

	status, body := controller.Update(arguments)

	if status != http.StatusBadRequest {
		suite.T().Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusBadRequest)
	}

	_, exists := body["error"]
	if !exists {
		suite.T().Errorf("An error is missing from the response body; it should've been there")
	}
}

func (suite *ConcealEmailControllerTestSuite) TestDescriptionUpdateFailedForUnkownReason() {
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{
			AddDescriptionReturnError: errors.New("Unknown error"),
		}
	})

	var arguments = map[string]interface{}{
		"concealEmailId": "an ID",
		"description":    "this value doesn't matter",
	}

	status, body := controller.Update(arguments)

	if status != http.StatusInternalServerError {
		suite.T().Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusInternalServerError)
	}

	_, exists := body["error"]
	if !exists {
		suite.T().Errorf("An error is missing from the response body; it should've been there")
	}
}

func (suite *ConcealEmailControllerTestSuite) TestDescriptionUpdateWithDelete() {

	concealEmailUsecase := TestConcealEmailUsecase{}
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &concealEmailUsecase
	})

	conceaEmailId := "an ID"
	var arguments = map[string]interface{}{
		"concealEmailId": conceaEmailId,
		"description":    "", //empty on purpose
	}

	status, body := controller.Update(arguments)

	if concealEmailUsecase.AddDescriptionReceiveConcealEmailPrefix != "" && concealEmailUsecase.DeleteDescriptionReceiveConcealEmailPrefix != conceaEmailId {
		suite.T().Errorf("The wrong usecase was called.  The delete usecase should have been called.")
	}

	if status != http.StatusOK {
		suite.T().Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusOK)
	}

	_, exists := body["error"]
	if exists {
		suite.T().Errorf("An error was returned in the response body; it shouldn't be there")
	}
}

func (suite *ConcealEmailControllerTestSuite) TestUpdateFailedWithConcealEmailNotExist() {

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
		"description":    "", //empty on purpose
	}

	status, body := controller.Update(arguments)

	if status != http.StatusNotFound {
		suite.T().Errorf("The returned status %d didn't equal the expected status of %d", status, http.StatusNotFound)
	}

	_, exists := body["error"]
	if !exists {
		suite.T().Errorf("An error is missing from the response body; it should've been there")
	}
}

//dependency injection mocks

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

//Start the test suite

func TestConcealEmailControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ConcealEmailControllerTestSuite))
}
