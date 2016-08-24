package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestRemoveAllTokens(c *C) {
	user, _, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})
	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	token, _, err := CreateToken(&CreateTokenRequest{
		UserID: user.ID,
	})
	c.Assert(err, IsNil)
	c.Check(token, Not(Equals), "")

	audit, err := RemoveAllTokens(&RemoveAllTokensRequest{})
	c.Assert(err, IsNil)
	c.Assert(audit, NotNil)

	tokens, _, err := GetTokens(&GetTokensRequest{
		UserID: user.ID,
	})
	c.Assert(err, IsNil)
	c.Check(tokens, DeepEquals, []string{})

	token, _, err = CreateToken(&CreateTokenRequest{
		UserID: user.ID,
	})
	c.Assert(err, IsNil)
	c.Check(token, Not(Equals), "")

	audit, err = RemoveAllTokens(&RemoveAllTokensRequest{
		UserID: user.ID,
	})
	c.Assert(err, IsNil)
	c.Assert(audit, NotNil)
	c.Check(audit.User, DeepEquals, user.Summary())

	tokens, _, err = GetTokens(&GetTokensRequest{
		UserID: user.ID,
	})
	c.Assert(err, IsNil)
	c.Check(tokens, DeepEquals, []string{})
}
