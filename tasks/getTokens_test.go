package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestGetTokens(c *C) {
	user, _, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})

	c.Assert(err, IsNil)

	token, _, err := CreateToken(&CreateTokenRequest{
		UserID: user.ID,
	})

	c.Assert(err, IsNil)

	tokens, audit, err := GetTokens(&GetTokensRequest{
		UserID: user.ID,
	})

	c.Assert(err, IsNil)
	c.Assert(audit, NotNil)
	c.Check(audit.User, DeepEquals, user.Summary())
	c.Assert(tokens, NotNil)
	c.Check(tokens, HasLen, 1)
	c.Check(tokens, DeepEquals, []string{token})
}
