package utils

import "net/http"

// Context represents an API request's context
type Context struct {
	Request         *http.Request
	Vars            map[string]string
	ResponseHeaders http.Header
	StatusCode      int
	User            User

	response http.ResponseWriter
}

// ReadBody will deserialize the request's body into the given object
func (c *Context) ReadBody(into interface{}) *Error {
	err := parseJSON(into, c)
	if err != nil {
		return NewError(400, "Bad Request", "The data you provided could not be parsed as valid JSON. Please check it and try again.")
	}

	return nil
}
