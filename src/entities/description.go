package entities

import (
	"fmt"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"unicode/utf8"
)

var minimumLength = 1
var maximumLength = 25
var DescriptionTooShortError = errors.New(fmt.Sprintf("The description must be at least %d characters long", minimumLength))
var DescriptionTooLongError = errors.New(fmt.Sprintf("The description can only be up to %d characters long", maximumLength))

func ValidateDescription(description string) error {
	length := utf8.RuneCountInString(description)

	if length < minimumLength {
		return DescriptionTooShortError
	} else if length > maximumLength {
		return DescriptionTooLongError
	}

	return nil
}
