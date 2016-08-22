package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestRemoveAllTokens(c *C) {
	user, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})
	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	token, err := CreateToken(&CreateTokenRequest{
		UserID: user.ID,
	})
	c.Assert(err, IsNil)
	c.Check(token, Not(Equals), "")

	err = RemoveAllTokens(&RemoveAllTokensRequest{})
	c.Assert(err, IsNil)

	tokens, err := GetTokens(&GetTokensRequest{
		UserID: user.ID,
	})
	c.Assert(err, IsNil)
	c.Check(tokens, DeepEquals, []string{})

	token, err = CreateToken(&CreateTokenRequest{
		UserID: user.ID,
	})
	c.Assert(err, IsNil)
	c.Check(token, Not(Equals), "")

	err = RemoveAllTokens(&RemoveAllTokensRequest{
		UserID: user.ID,
	})
	c.Assert(err, IsNil)

	tokens, err = GetTokens(&GetTokensRequest{
		UserID: user.ID,
	})
	c.Assert(err, IsNil)
	c.Check(tokens, DeepEquals, []string{})
}
