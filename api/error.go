package api

import "github.com/pkg/errors"

type Error struct {
	Err        error  // Original error stack for logging.
	Message    string // Error message for api response.
	StatusCode int    // Related error's http status code.
}

func (e *Error) Error() string {
	return e.Err.Error()
}

// Create new API error object with original error stack, supposed response message and status code.
func NewErrorWrap(err error, stack, message string, code int) *Error {
	return &Error{
		Err:        errors.Wrap(err, stack),
		Message:    message,
		StatusCode: code,
	}
}
