package errs

import "fmt"

// Error is the error type.
type Error struct {
	Code    int32
	Message string
}

const (
	// ErrCodeUnknown is the unknown error code.
	ErrCodeUnknown = 1000

	// ErrConfigInit is the error code for configuration initialization.
	ErrConfigInit = 2001
	// ErrDBConnect is the error code for database connection.
	ErrDBConnect = 2002
)

var errMp = map[int32]string{
	ErrCodeUnknown: "unknown error",
}

// New creates a new error.
func New(code int32, message string) error {
	return Error{Code: code, Message: message}
}

// Code creates a new error with the given code.
func Code(code int32) error {
	if message, ok := errMp[code]; ok {
		return New(code, message)
	}
	return New(code, errMp[ErrCodeUnknown])
}

// Msg creates a new error with the given message.
func Msg(message string) error {
	return New(ErrCodeUnknown, message)
}

// Error returns the error message.
func (e Error) Error() string {
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}
