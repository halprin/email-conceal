package errors

import "github.com/pkg/errors"

func New(message string) error {
	return errors.New(message)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
