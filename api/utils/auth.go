package utils

// MustBeAuthenticated checks that the user has been authenticated and returns
// an API compatible error if they are not.
func (c *Context) MustBeAuthenticated() *Error {
	if !c.IsAuthenticated() {
		return NewError(401, "Unauthorized", "You have not provided a valid authorization token with your request.")
	}

	return nil
}

// IsAuthenticated asserts that the request has been authenticated
func (c *Context) IsAuthenticated() bool {
	return c.User != nil
}

// RequireAuthentication configures this handler to require authentication
func (h *Handler) RequireAuthentication() *Handler {
	return h.RegisterPreprocessors(func(c *Context) *Error {
		return c.MustBeAuthenticated()
	})
}
