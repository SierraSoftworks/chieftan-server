package utils

import . "gopkg.in/check.v1"

type TestUser struct {
	id          string
	permissions []string
}

func (u *TestUser) ID() string {
	return u.id
}

func (u *TestUser) Permissions() []string {
	return u.permissions
}

func (s *TestSuite) TestIsAuthenticated(c *C) {
	ctx := &Context{}
	c.Check(ctx.IsAuthenticated(), Equals, false)

	ctx = &Context{
		User: &TestUser{
			id:          "bob",
			permissions: []string{"x", "y", "z"},
		},
	}

	c.Check(ctx.IsAuthenticated(), Equals, true)
}

func (s *TestSuite) TestMustBeAuthenticated(c *C) {
	ctx := &Context{}
	c.Check(ctx.MustBeAuthenticated(), NotNil)
	c.Check(ctx.MustBeAuthenticated(), FitsTypeOf, &Error{})

	ctx = &Context{
		User: &TestUser{
			id:          "bob",
			permissions: []string{"x", "y", "z"},
		},
	}

	c.Check(ctx.MustBeAuthenticated(), IsNil)
}
