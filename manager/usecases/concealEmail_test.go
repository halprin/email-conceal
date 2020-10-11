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
