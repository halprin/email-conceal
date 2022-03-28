package concealEmail

import (
	"errors"
	"fmt"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/entities"
	"github.com/stretchr/testify/suite"
	"testing"
)

var usecase = ConcealEmailUsecaseImpl{}
var testAppContext = context.ApplicationContext{}

type ConcealEmailTestSuite struct {
	suite.Suite
}

func (suite *ConcealEmailTestSuite) SetupTest() {
	//defaults each test to a dummy set of values
	testAppContext.Reset()
	testAppContext.Bind(func() ConcealEmailGateway {
		return &TestConcealEmailGateway{} //no good default for this, so nothing set
	})
	testAppContext.Bind(func() context.EnvironmentGateway {
		return &TestEnvironmentGateway{
			GetReturnMap: map[string]string{
				"DOMAIN": "example.com",
			},
		}
	})
	testAppContext.Bind(func() context.UuidLibrary {
		return &TestUuidLibrary{
			GenerateReturnUuid: "random-uuid",
		}
	})
	usecase.Init()
}

func (suite *ConcealEmailTestSuite) TestAddConcealEmailSuccess() {
	uuid := "moof-uuid"
	domain := "dogcow.com"

	testConcealGateway := TestConcealEmailGateway{
		GetActualEmailDetails_ReturnIsVerified: true,
	}
	testAppContext.Bind(func() ConcealEmailGateway {
		return &testConcealGateway
	})
	testEnvironmentGateway := TestEnvironmentGateway{
		GetReturnMap: map[string]string{
			"DOMAIN": domain,
		},
	}
	testAppContext.Bind(func() context.EnvironmentGateway {
		return &testEnvironmentGateway
	})
	testUuidLibrary := TestUuidLibrary{
		GenerateReturnUuid: uuid,
	}
	testAppContext.Bind(func() context.UuidLibrary {
		return &testUuidLibrary
	})
	usecase.Init()

	description := "description"

	actualConcealedEmail, err := usecase.Add("valid-email@dogcow.com", &description)

	suite.Assert().Nil(err)

	expectedConcealedEmail := fmt.Sprintf("%s@%s", uuid, domain)
	suite.Assert().Equal(expectedConcealedEmail, actualConcealedEmail)
}

func (suite *ConcealEmailTestSuite) TestAddConcealEmailSuccessWithNoDescription() {
	uuid := "moof-uuid"
	domain := "dogcow.com"

	testConcealGateway := TestConcealEmailGateway{
		GetActualEmailDetails_ReturnIsVerified: true,
	}
	testAppContext.Bind(func() ConcealEmailGateway {
		return &testConcealGateway
	})
	testEnvironmentGateway := TestEnvironmentGateway{
		GetReturnMap: map[string]string{
			"DOMAIN": domain,
		},
	}
	testAppContext.Bind(func() context.EnvironmentGateway {
		return &testEnvironmentGateway
	})
	testUuidLibrary := TestUuidLibrary{
		GenerateReturnUuid: uuid,
	}
	testAppContext.Bind(func() context.UuidLibrary {
		return &testUuidLibrary
	})
	usecase.Init()

	actualConcealedEmail, err := usecase.Add("valid-email@dogcow.com", nil)

	suite.Assert().Nil(err)

	expectedConcealedEmail := fmt.Sprintf("%s@%s", uuid, domain)
	suite.Assert().Equal(expectedConcealedEmail, actualConcealedEmail)
}

func (suite *ConcealEmailTestSuite) TestAddConcealFailedForBadEmail() {

	description := "description"

	_, err := usecase.Add("in[valid-email@dogcow.com", &description)

	suite.Assert().ErrorIs(err, entities.InvalidEmailAddressError)
}

func (suite *ConcealEmailTestSuite) TestFailedToAddTheMapping() {
	expectedError := errors.New("oops")
	testGateway := TestConcealEmailGateway{
		GetActualEmailDetails_ReturnIsVerified: true,
		AddReturnError:                         expectedError,
	}
	testAppContext.Bind(func() ConcealEmailGateway {
		return &testGateway
	})
	usecase.Init()

	description := "description"

	_, err := usecase.Add("moof@dogcow.com", &description)

	suite.Assert().ErrorIs(err, expectedError)
}

func (suite *ConcealEmailTestSuite) TestConcealEmailBadDescription() {

	testGateway := TestConcealEmailGateway{
		GetActualEmailDetails_ReturnIsVerified: true,
		AddReturnError:                         errors.New("oops"),
	}
	testAppContext.Bind(func() ConcealEmailGateway {
		return &testGateway
	})
	usecase.Init()

	description := ""

	_, err := usecase.Add("moof@dogcow.com", &description)

	suite.Assert().ErrorIs(err, entities.DescriptionTooShortError)
}

func (suite *ConcealEmailTestSuite) TestDeleteConcealEmailSuccess() {
	err := usecase.Delete("some_prefix")

	suite.Assert().Nil(err)
}

func (suite *ConcealEmailTestSuite) TestDeleteConcealEmailNegative() {
	expectedError := errors.New("it failed")
	testGateway := TestConcealEmailGateway{
		DeleteReturnError: expectedError,
	}
	testAppContext.Bind(func() ConcealEmailGateway {
		return &testGateway
	})
	usecase.Init()

	err := usecase.Delete("some_prefix")

	suite.Assert().ErrorIs(err, expectedError)
}

func (suite *ConcealEmailTestSuite) TestAddDescriptionFailsForEntityError() {

	err := usecase.AddDescriptionToExistingEmail("some_prefix", "")

	suite.Assert().ErrorIs(err, entities.DescriptionTooShortError)
}

func (suite *ConcealEmailTestSuite) TestAddDescriptionFailsForGatewayFailure() {
	expectedError := errors.New("an error")
	testAppContext.Bind(func() ConcealEmailGateway {
		return &TestConcealEmailGateway{
			UpdateReturnError: expectedError,
		}
	})
	usecase.Init()

	err := usecase.AddDescriptionToExistingEmail("some_prefix", "a description")

	suite.Assert().ErrorIs(err, expectedError)
}

func (suite *ConcealEmailTestSuite) TestAddDescriptionSuccess() {
	testGateway := TestConcealEmailGateway{}
	testAppContext.Bind(func() ConcealEmailGateway {
		return &testGateway
	})
	usecase.Init()

	prefix := "some_prefix"
	description := "a description"
	err := usecase.AddDescriptionToExistingEmail(prefix, description)

	suite.Assert().Nil(err)

	suite.Assert().Equal(prefix, testGateway.UpdateReceiveConcealPrefix)

	suite.Assert().Equal(description, *testGateway.UpdateReceiveDescription)
}

func (suite *ConcealEmailTestSuite) TestDeleteDescriptionFailed() {
	expectedError := errors.New("an error")
	testAppContext.Bind(func() ConcealEmailGateway {
		return &TestConcealEmailGateway{
			UpdateReturnError: expectedError,
		}
	})
	usecase.Init()

	err := usecase.DeleteDescriptionFromExistingEmail("some_prefix")

	suite.Assert().ErrorIs(err, expectedError)
}

func (suite *ConcealEmailTestSuite) TestDeleteDescriptionSuccess() {
	testGateway := TestConcealEmailGateway{}
	testAppContext.Bind(func() ConcealEmailGateway {
		return &testGateway
	})
	usecase.Init()

	prefix := "some_prefix"
	err := usecase.DeleteDescriptionFromExistingEmail(prefix)

	suite.Assert().Nil(err)

	suite.Assert().Equal(prefix, testGateway.UpdateReceiveConcealPrefix)

	suite.Assert().Nil(testGateway.UpdateReceiveDescription)
}

//dependency injection mocks

type TestConcealEmailGateway struct {
	AddReceiveConcealPrefix string
	AddReceiveActualEmail   string
	AddReceiveDescription   *string
	AddReturnError          error

	DeleteReceiveConcealPrefix string
	DeleteReturnError          error

	UpdateReceiveConcealPrefix string
	UpdateReceiveDescription   *string
	UpdateReturnError          error

	GetActualEmailDetails_ActualEmail        string
	GetActualEmailDetails_ReturnEmailAddress string
	GetActualEmailDetails_ReturnIsVerified   bool
	GetActualEmailDetails_ReturnError        error
}

func (testGateway *TestConcealEmailGateway) AddConcealedEmailToActualEmailMapping(concealPrefix string, actualEmail string, description *string) error {
	testGateway.AddReceiveConcealPrefix = concealPrefix
	testGateway.AddReceiveActualEmail = actualEmail
	testGateway.AddReceiveDescription = description

	return testGateway.AddReturnError
}

func (testGateway *TestConcealEmailGateway) DeleteConcealedEmailToActualEmailMapping(concealPrefix string) error {
	testGateway.DeleteReceiveConcealPrefix = concealPrefix
	return testGateway.DeleteReturnError
}

func (testGateway *TestConcealEmailGateway) UpdateConcealedEmail(concealPrefix string, description *string) error {
	testGateway.UpdateReceiveConcealPrefix = concealPrefix
	testGateway.UpdateReceiveDescription = description
	return testGateway.UpdateReturnError
}

func (testGateway *TestConcealEmailGateway) GetActualEmailDetails(actualEmail string) (string, bool, error) {
	testGateway.GetActualEmailDetails_ActualEmail = actualEmail
	return testGateway.GetActualEmailDetails_ReturnEmailAddress, testGateway.GetActualEmailDetails_ReturnIsVerified, testGateway.GetActualEmailDetails_ReturnError
}

type TestEnvironmentGateway struct {
	GetReceiveKey string
	GetReturnMap  map[string]string
}

func (testGateway *TestEnvironmentGateway) GetEnvironmentValue(key string) string {
	testGateway.GetReceiveKey = key
	return testGateway.GetReturnMap[key]
}

type TestUuidLibrary struct {
	GenerateReturnUuid string
}

func (testLibrary *TestUuidLibrary) GenerateRandomUuid() string {
	return testLibrary.GenerateReturnUuid
}

//Start the test suite

func TestConcealEmailTestSuite(t *testing.T) {
	suite.Run(t, new(ConcealEmailTestSuite))
}
