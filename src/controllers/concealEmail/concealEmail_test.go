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
	//defaults each test to a dummy set of values
	testAppContext.Reset()
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{}
	})
	controller.Init()
}

func (suite *ConcealEmailControllerTestSuite) TestConcealEmailControllerSuccess() {
	concealedEmail := "concealed@asdf.com"

	testUsecase := TestConcealEmailUsecase{
		AddReturnConcealEmail: concealedEmail,
	}
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &testUsecase
	})
	controller.Init()

	sourceEmail := "dogcow@apple.com"
	var arguments = map[string]interface{}{
		"email": sourceEmail,
	}

	status, body := controller.Add(arguments)
	actualConcealedEmail := body["concealedEmail"]

	suite.Assert().Equal(sourceEmail, testUsecase.AddReceiveSourceEmail)
	suite.Assert().Equal(concealedEmail, actualConcealedEmail)
	suite.Assert().Equal(http.StatusCreated, status)
}

func (suite *ConcealEmailControllerTestSuite) TestConcealEmailControllerBadEmailType() {

	testUsecase := TestConcealEmailUsecase{}
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &testUsecase
	})
	controller.Init()

	var arguments = map[string]interface{}{
		"email": 3,
	}

	status, _ := controller.Add(arguments)

	suite.Assert().Empty(testUsecase.AddReceiveSourceEmail)
	suite.Assert().Equal(http.StatusBadRequest, status)
}

func (suite *ConcealEmailControllerTestSuite) TestConcealEmailControllerInvalidEmail() {

	testUsecase := TestConcealEmailUsecase{
		AddReturnError: entities.InvalidEmailAddressError,
	}
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &testUsecase
	})
	controller.Init()

	sourceEmail := "dogcow"
	var arguments = map[string]interface{}{
		"email": sourceEmail,
	}

	status, _ := controller.Add(arguments)

	suite.Assert().Equal(sourceEmail, testUsecase.AddReceiveSourceEmail)
	suite.Assert().Equal(http.StatusBadRequest, status)
}

func (suite *ConcealEmailControllerTestSuite) TestConcealEmailControllerUnknownError() {

	testUsecase := TestConcealEmailUsecase{
		AddReturnError: errors.New("some other error"),
	}
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &testUsecase
	})
	controller.Init()

	sourceEmail := "dogcow@apple.com"
	var arguments = map[string]interface{}{
		"email": sourceEmail,
	}

	status, _ := controller.Add(arguments)

	suite.Assert().Equal(sourceEmail, testUsecase.AddReceiveSourceEmail)
	suite.Assert().Equal(http.StatusInternalServerError, status)
}

func (suite *ConcealEmailControllerTestSuite) TestDeleteConcealEmailControllerSuccess() {

	var arguments = map[string]interface{}{
		"concealEmailId": "dogcow",
	}

	status, body := controller.Delete(arguments)

	suite.Assert().Equal(http.StatusNoContent, status)
	suite.Assert().Empty(body)
}

func (suite *ConcealEmailControllerTestSuite) TestDeleteConcealEmailControllerBadInput() {

	var arguments = map[string]interface{}{
		"concealEmailId": 3,
	}

	status, body := controller.Delete(arguments)

	suite.Assert().Equal(http.StatusBadRequest, status)

	_, exists := body["error"]
	suite.Assert().True(exists)
}

func (suite *ConcealEmailControllerTestSuite) TestDeleteConcealEmailControllerFailedDelete() {

	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{
			DeleteReturnError: errors.New("moof! go boom"),
		}
	})
	controller.Init()

	var arguments = map[string]interface{}{
		"concealEmailId": "dogcow",
	}

	status, body := controller.Delete(arguments)

	suite.Assert().Equal(http.StatusInternalServerError, status)

	_, exists := body["error"]
	suite.Assert().True(exists)
}

func (suite *ConcealEmailControllerTestSuite) TestUpdateConcealEmailWithNewDescription() {

	var arguments = map[string]interface{}{
		"concealEmailId": "an ID",
		"description":    "a new description",
	}

	status, body := controller.Update(arguments)

	suite.Assert().Equal(http.StatusOK, status)
	suite.Assert().Empty(body)
}

func (suite *ConcealEmailControllerTestSuite) TestTooLongDescriptionUpdate() {
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{
			AddDescriptionReturnError: entities.DescriptionTooLongError,
		}
	})
	controller.Init()

	var arguments = map[string]interface{}{
		"concealEmailId": "an ID",
		"description":    "this value doesn't matter",
	}

	status, body := controller.Update(arguments)

	suite.Assert().Equal(http.StatusBadRequest, status)

	_, exists := body["error"]
	suite.Assert().True(exists)
}

func (suite *ConcealEmailControllerTestSuite) TestDescriptionUpdateFailedForUnkownReason() {
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &TestConcealEmailUsecase{
			AddDescriptionReturnError: errors.New("Unknown error"),
		}
	})
	controller.Init()

	var arguments = map[string]interface{}{
		"concealEmailId": "an ID",
		"description":    "this value doesn't matter",
	}

	status, body := controller.Update(arguments)

	suite.Assert().Equal(http.StatusInternalServerError, status)

	_, exists := body["error"]
	suite.Assert().True(exists)
}

func (suite *ConcealEmailControllerTestSuite) TestDescriptionUpdateWithDelete() {

	testUsecase := TestConcealEmailUsecase{}
	testAppContext.Bind(func() concealEmail.ConcealEmailUsecase {
		return &testUsecase
	})
	controller.Init()

	conceaEmailId := "an ID"
	var arguments = map[string]interface{}{
		"concealEmailId": conceaEmailId,
		"description":    "", //empty on purpose
	}

	status, body := controller.Update(arguments)

	//tests that the correct usecase was called
	suite.Assert().Empty(testUsecase.AddDescriptionReceiveConcealEmailPrefix)
	suite.Assert().Equal(conceaEmailId, testUsecase.DeleteDescriptionReceiveConcealEmailPrefix)

	suite.Assert().Equal(http.StatusOK, status)

	_, exists := body["error"]
	suite.Assert().False(exists)
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
	controller.Init()

	var arguments = map[string]interface{}{
		"concealEmailId": conceaEmailId,
		"description":    "", //empty on purpose
	}

	status, body := controller.Update(arguments)

	suite.Assert().Equal(http.StatusNotFound, status)

	_, exists := body["error"]
	suite.Assert().True(exists)
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
