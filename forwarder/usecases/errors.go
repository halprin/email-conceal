package usecases

import (
	"fmt"
	"io"
	"strings"
)

//top level error for the forwarder
type ForwarderError struct {
	cause   error
	message string
}

func (error *ForwarderError) Error() string {
	errorMessageBuilder := strings.Builder{}

	errorMessageBuilder.WriteString(error.message)
	errorMessageBuilder.WriteString(": ")
	errorMessageBuilder.WriteString(error.Unwrap().Error())

	return errorMessageBuilder.String()
}

func (error *ForwarderError) Unwrap() error {
	return error.cause
}

func (error *ForwarderError) Cause() error {
	return error.cause
}

func (error *ForwarderError) Is(target error) bool {
	_, ok := target.(*ForwarderError)
	return ok
}

func (error *ForwarderError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = io.WriteString(s, error.message)
			_, _ = io.WriteString(s, "\nCaused by...\n")
			_, _ = fmt.Fprintf(s, "%+v", error.Unwrap())
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, error.Error())
	}
}

func (error *ForwarderError) SetMessage(message string) {
	error.message = message
}

func (error *ForwarderError) SetCause(cause error) {
	error.cause = cause
}

//sub error
type UnableToReadEmailError struct {
	ForwarderError
}

func (error *UnableToReadEmailError) Is(target error) bool {
	_, ok := target.(*UnableToReadEmailError)
	return ok
}

func NewUnableToReadEmailError(url string, cause error) *UnableToReadEmailError {
	err := &UnableToReadEmailError{}

	err.SetMessage(fmt.Sprintf("Unable to read e-mail at %s", url))
	err.SetCause(cause)

	return err
}

//sub error
type UnableToParseEmailError struct {
	ForwarderError
}

func (error *UnableToParseEmailError) Is(target error) bool {
	_, ok := target.(*UnableToParseEmailError)
	return ok
}

func NewUnableToParseEmailError(cause error) error {
	err := &UnableToParseEmailError{}

	err.SetMessage("Unable to parse raw e-mail")
	err.SetCause(cause)

	return err
}

//sub error
type UnableToSendEmailError struct {
	ForwarderError
}

func (error *UnableToSendEmailError) Is(target error) bool {
	_, ok := target.(*UnableToSendEmailError)
	return ok
}

func NewUnableToSendEmailError(cause error) error {
	err := &UnableToSendEmailError{}

	err.SetMessage("Unable to send e-mail")
	err.SetCause(cause)

	return err
}
