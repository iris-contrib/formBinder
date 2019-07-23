package formbinder

import (
	"encoding/json"
	"fmt"
)

// Error wraps an error.
type Error struct {
	err error
}

func (s *Error) Error() string {
	return s.err.Error()
}

// MarshalJSON completes the json.Marshaller.
func (s Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.err.Error())
}

// Cause implements the causer interface from github.com/pkg/errors.
func (s *Error) Cause() error {
	return s.err
}

func newError(format string, a ...interface{}) error {
	return &Error{fmt.Errorf(format, a...)}
}

// IsErrPath reports whether the incoming error is type of `ErrPath`, which can be ignored
// when server allows unknown post values to be sent by the client.
func IsErrPath(err error) bool {
	if err == nil {
		return false
	}

	_, ok := err.(ErrPath)
	return ok
}

// ErrPath describes an error that can be ignored if server allows unknown post values to be sent on server.
type ErrPath struct {
	field string
}

func (err ErrPath) Error() string {
	return fmt.Sprintf("form: not found the field \"%s\"", err.field)
}

// Field returns the unknown posted request field.
func (err ErrPath) Field() string {
	return err.field
}
