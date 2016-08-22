package utils

// Preprocessor entries are executed during the request's pipeline and
// can pre-emptively kick out of the processing stage by raising
// an error.
type Preprocessor func(c *Context) *Error
