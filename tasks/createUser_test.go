package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestCreateUser(c *C) {
	req := &CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
		Permissions: []string{
			"admin",
		},
	}

	user, audit, err := CreateUser(req)
	c.Assert(err, IsNil)
	c.Assert(audit, NotNil)
	c.Check(audit.User, NotNil)
	c.Check(audit.User.ID, Equals, user.ID)

	c.Assert(user, NotNil)
	c.Check(user.ID, Equals, "b642b4217b34b1e8d3bd915fc65c4452")
	c.Check(user.Name, Equals, "Test User")
	c.Check(user.Email, Equals, "test@test.com")
	c.Check(user.Permissions, DeepEquals, []string{"admin"})
	c.Check(user.Tokens, DeepEquals, []string{})
}
