package entities

import (
	"errors"
	"testing"
)

func TestValidateDescriptionTooShort(t *testing.T) {
	err := ValidateDescription("")

	if !errors.Is(err, DescriptionTooShortError) {
		t.Error("Expected the description to fail the too short validation")
	}
}

func TestValidateDescriptionTooLong(t *testing.T) {
	description := string(make([]rune, maximumLength + 1))
	err := ValidateDescription(description)

	if !errors.Is(err, DescriptionTooLongError) {
		t.Error("Expected the description to fail the too short validation")
	}
}
