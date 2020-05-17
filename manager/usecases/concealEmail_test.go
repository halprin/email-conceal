package usecases

import (
	"fmt"
	"github.com/halprin/email-conceal/manager/context"
	"testing"
)

func TestConcealEmail(t *testing.T) {
	uuid := "moof-uuid"
	domain := "dogcow.com"
	testApplicationContext := &context.TestApplicationContext{
		ReturnFromGenerateRandomUuid: uuid,
		ReturnFromEnvironmentGateway: map[string]string{
			"DOMAIN": domain,
		},
	}

	actualConcealedEmail, err := AddConcealEmailUsecase("valid-email@dogcow.com", testApplicationContext)

	if err != nil {
		t.Error("Expected no error to be returned from concealing the e-mail usecase, but there was one")
	}

	expectedConcealedEmail := fmt.Sprintf("%s@%s", uuid, domain)
	if actualConcealedEmail != expectedConcealedEmail {
		t.Errorf("The generated concealed e-mail %s was supposed to be %s", actualConcealedEmail, expectedConcealedEmail)
	}
}

func TestConcealEmailNegative(t *testing.T) {
	testApplicationContext := &context.TestApplicationContext{}

	_, err := AddConcealEmailUsecase("in[valid-email@dogcow.com", testApplicationContext)

	if err == nil {
		t.Error("Expected an error to be returned from concealing the e-mail usecase, but there wasn't one")
	}
}
