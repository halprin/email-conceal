package usecases

import (
	"errors"
	"fmt"
	"github.com/halprin/email-conceal/manager/context"
	"testing"
)


var usecase = ConcealEmailUsecaseImpl{}
var testAppContext = context.ApplicationContext{}

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

func TestConcealEmailSuccess(t *testing.T) {
	uuid := "moof-uuid"
	domain := "dogcow.com"

	testConcealGateway := TestConcealEmailGateway{}
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

	description := "description"

	actualConcealedEmail, err := usecase.Add("valid-email@dogcow.com", &description)

	if err != nil {
		t.Error("Expected no error to be returned from concealing the e-mail usecase, but there was one")
	}

	expectedConcealedEmail := fmt.Sprintf("%s@%s", uuid, domain)
	if actualConcealedEmail != expectedConcealedEmail {
		t.Errorf("The generated concealed e-mail %s was supposed to be %s", actualConcealedEmail, expectedConcealedEmail)
	}
}

func TestConcealEmailSuccessWithNoDescription(t *testing.T) {
	uuid := "moof-uuid"
	domain := "dogcow.com"

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

	actualConcealedEmail, err := usecase.Add("valid-email@dogcow.com", nil)

	if err != nil {
		t.Error("Expected no error to be returned from concealing the e-mail usecase, but there was one")
	}

	expectedConcealedEmail := fmt.Sprintf("%s@%s", uuid, domain)
	if actualConcealedEmail != expectedConcealedEmail {
		t.Errorf("The generated concealed e-mail %s was supposed to be %s", actualConcealedEmail, expectedConcealedEmail)
	}
}

func TestConcealEmailBadEmail(t *testing.T) {
	description := "description"

	_, err := usecase.Add("in[valid-email@dogcow.com", &description)

	if err == nil {
		t.Error("Expected an error to be returned from concealing the e-mail usecase, but there wasn't one")
	}
}

func TestConcealEmailGatewayFailed(t *testing.T) {

	testGateway := TestConcealEmailGateway{
		AddReturnError: errors.New("oops"),
	}
	testAppContext.Bind(func() ConcealEmailGateway {
		return &testGateway
	})

	description := "description"

	_, err := usecase.Add("moof@dogcow.com", &description)

	if err == nil {
		t.Error("Expected an error to be returned from concealing the e-mail usecase, but there wasn't one")
	}
}

func TestConcealEmailBadDescription(t *testing.T) {
	testGateway := TestConcealEmailGateway{
		AddReturnError: errors.New("oops"),
	}
	testAppContext.Bind(func() ConcealEmailGateway {
		return &testGateway
	})

	description := ""

	_, err := usecase.Add("moof@dogcow.com", &description)

	if err == nil {
		t.Error("Expected an error to be returned from concealing the e-mail usecase, but there wasn't one")
	}
}

func TestDeleteConcealEmailSuccess(t *testing.T) {

	err := usecase.Delete("some_prefix")

	if err != nil {
		t.Error("Expected no error to be returned from the delete conceal usecase, but there was one")
	}
}

func TestDeleteConcealEmailNegative(t *testing.T) {

	testGateway := TestConcealEmailGateway{
		DeleteReturnError: errors.New("it failed"),
	}
	testAppContext.Bind(func() ConcealEmailGateway {
		return &testGateway
	})

	err := usecase.Delete("some_prefix")

	if err == nil {
		t.Error("Expected an error to be returned from the delete conceal usecase, but there wasn't one")
	}
}

func TestAddDescriptionFailsForEntityError(t *testing.T) {

	err := usecase.AddDescriptionToExistingEmail("some_prefix", "")

	if err == nil {
		t.Error("Expected an error to be returned from the update conceal usecase, but there wasn't one")
	}
}

func TestAddDescriptionFailsForGatewayFailure(t *testing.T) {

	testAppContext.Bind(func() ConcealEmailGateway {
		return &TestConcealEmailGateway{
			UpdateReturnError: errors.New("an error"),
		}
	})

	err := usecase.AddDescriptionToExistingEmail("some_prefix", "a description")

	if err == nil {
		t.Error("Expected an error to be returned from the update conceal usecase, but there wasn't one")
	}
}

func TestAddDescriptionSuccess(t *testing.T) {

	testGateway := TestConcealEmailGateway{}
	testAppContext.Bind(func() ConcealEmailGateway {
		return &testGateway
	})

	prefix := "some_prefix"
	description := "a description"
	err := usecase.AddDescriptionToExistingEmail(prefix, description)

	if err != nil {
		t.Error("An error was returned from the add description usecase, but it wasn't expected")
	}

	if testGateway.UpdateReceiveConcealPrefix != prefix {
		t.Errorf("The update gateway wasn't called with the prefix %s, instead it was called with %s", prefix, testGateway.UpdateReceiveConcealPrefix)
	}

	if *testGateway.UpdateReceiveDescription != description {
		t.Errorf("The update gateway wasn't called with the description %s, instead it was called with %s", description, *testGateway.UpdateReceiveDescription)
	}
}

func TestDeleteDescriptionFailed(t *testing.T) {
	testAppContext.Bind(func() ConcealEmailGateway {
		return &TestConcealEmailGateway{
			UpdateReturnError: errors.New("an error"),
		}
	})

	err := usecase.DeleteDescriptionFromExistingEmail("some_prefix")

	if err == nil {
		t.Error("An error wasn't returned from the delete description usecase, but it was supposed to")
	}
}

func TestDeleteDescriptionSuccess(t *testing.T) {
	testGateway := TestConcealEmailGateway{}
	testAppContext.Bind(func() ConcealEmailGateway {
		return &testGateway
	})

	prefix := "some_prefix"
	err := usecase.DeleteDescriptionFromExistingEmail(prefix)

	if err != nil {
		t.Error("An error was returned from the delete description usecase, but it wasn't expected")
	}

	if testGateway.UpdateReceiveConcealPrefix != prefix {
		t.Errorf("The update gateway wasn't called with the prefix %s, instead it was called with %s", prefix, testGateway.UpdateReceiveConcealPrefix)
	}

	if testGateway.UpdateReceiveDescription != nil {
		t.Errorf("The update gateway wasn't called with a nil description, instead it was called with %s", *testGateway.UpdateReceiveDescription)
	}
}
