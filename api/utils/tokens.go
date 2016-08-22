package utils

import "strings"

// AuthorizationToken represents an authorization token provided in the Authorization
// header of a request.
type AuthorizationToken struct {
	Type  string
	Value string
}

// GetAuthToken will retrieve the authorization token provided with the request
// if it exists and is syntactically valid, otherwise it'll return null.
func (c *Context) GetAuthToken() *AuthorizationToken {
	authHdrs := c.Request.Header["Authorization"]
	if len(authHdrs) != 1 {
		return nil
	}

	parts := strings.SplitN(authHdrs[0], " ", 2)
	if len(parts) < 2 {
		return nil
	}

	return &AuthorizationToken{
		Type:  parts[0],
		Value: parts[1],
	}
}
