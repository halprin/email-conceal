package entities

import (
	"github.com/halprin/email-conceal/external/lib/errors"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	validEmail := "dogcow@apple.com"

	err := ValidateEmail(validEmail)

	if err != nil && errors.Is(err, InvalidEmailAddressError) {
		t.Errorf("%s is a valid e-mail, but test thinks it isn't: %+v", validEmail, err)
	} else if err != nil {
		t.Errorf("Some other error occured that wasn't an InvalidEmailAddressError: %+v", err)
	}
}

func TestValidateEmailNegative(t *testing.T) {
	invalidEmail := "dog[cow@apple.com"

	err := ValidateEmail(invalidEmail)

	if err == nil {
		t.Errorf("%s is an invalid e-mail, but test thinks it valid", invalidEmail)
	} else if !errors.Is(err, InvalidEmailAddressError) {
		t.Errorf("Some other error occured that wasn't an InvalidEmailAddressError: %+v", err)
	}
}
