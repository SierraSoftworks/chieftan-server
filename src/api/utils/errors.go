package utils

import "fmt"

// Error represents an API level error object built up of a code, name and message
type Error struct {
	Code    int    `json:"code"`
	Name    string `json:"error"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d %s: %s", e.Code, e.Name, e.Message)
}

// NewError creates a new API level error object with the given properties
func NewError(code int, name, message string) *Error {
	return &Error{
		Code:    code,
		Name:    name,
		Message: message,
	}
}

// NewErrorFor creates a new API level error object for a general error object
func NewErrorFor(err error) *Error {
	return &Error{
		Code:    500,
		Name:    "Server Error",
		Message: "An unexpected error occured while processing your request. Please check it and try again.",
	}
}
