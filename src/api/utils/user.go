package utils

// User represents an authenticated user
type User interface {
	ID() string
	Permissions() []string
}

// UserStore represents something capable of providing users
// on request based on their authorization tokens.
type UserStore interface {
	GetUser(token *AuthorizationToken) (User, *Error)
}

// ActiveUserStore should be set to a UserStore implementation
// which will provide users whenever an authorization token is
// present in a request.
var ActiveUserStore UserStore
