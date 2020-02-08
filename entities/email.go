package entities

import (
	"github.com/halprin/email-conceal/external/lib/errors"
	"regexp"
)

//https://html.spec.whatwg.org/multipage/input.html#valid-e-mail-address
const emailRegex = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
var compiledEmailRegex *regexp.Regexp
var InvalidEmailAddressError = errors.New("E-mail address is invalid")

func ValidateEmail(email string) error {
	if compiledEmailRegex == nil {
		var err error
		compiledEmailRegex, err = regexp.Compile(emailRegex)
		if err != nil {
			return errors.Wrap(err, "Compiling the e-mail regex failed")
		}
	}

	emailIsValid := compiledEmailRegex.MatchString(email)

	if emailIsValid == false {
		return InvalidEmailAddressError
	}

	return nil
}
