package utils

import "regexp"

// MustHavePermission checks that the user has the provided permission and
// returns an API compatible error if they do not.
func (c *Context) MustHavePermission(permission string) *Error {
	if !c.HasPermission(permission) {
		return NewError(403, "Not Allowed", "You do not have permission to make use of this method.")
	}

	return nil
}

// RequirePermissions configures this handler to require the provided permissions for all requests
func (h *Handler) RequirePermissions(permissions ...string) *Handler {
	h.RequireAuthentication()

	for _, p := range permissions {
		h.RegisterPreprocessors(func(c *Context) *Error {
			return c.MustHavePermission(p)
		})
	}

	return h
}

// HasPermission asserts that the current user is in posession of the specified permission
func (c *Context) HasPermission(permission string) bool {
	if c.User == nil {
		return false
	}

	userPermissions := c.User.Permissions()

	for _, p := range userPermissions {
		if p == permission {
			return true
		}
	}

	permRx := regexp.MustCompile("\\:\\w+")
	actualPermission := permRx.ReplaceAllStringFunc(permission, func(match string) string {
		name := match[1:]
		replacement, hasReplacement := c.Vars[name]
		if !hasReplacement {
			return match
		}

		return replacement
	})

	for _, p := range userPermissions {
		if p == actualPermission {
			return true
		}
	}

	return false
}
