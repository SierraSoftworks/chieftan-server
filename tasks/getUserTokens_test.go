package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestGetUserTokens(c *C) {
	newUser, _, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})

	c.Assert(err, IsNil)

	token, _, err := CreateToken(&CreateTokenRequest{
		UserID: newUser.ID,
	})

	tokens, audit, err := GetUserTokens(&GetUserTokensRequest{
		ID: newUser.ID,
	})

	c.Assert(err, IsNil)
	c.Assert(audit, NotNil)
	c.Assert(audit.User, NotNil)
	c.Check(audit.User.Name, Equals, "Test User")
	c.Check(audit.User.Email, Equals, "test@test.com")
	c.Assert(tokens, NotNil)
	c.Check(tokens, HasLen, 1)
	c.Check(tokens, DeepEquals, []string{token})
}
