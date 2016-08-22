package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestGetUserTokens(c *C) {
	newUser, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})

	c.Assert(err, IsNil)

	token, err := CreateToken(&CreateTokenRequest{
		UserID: newUser.ID,
	})

	tokens, err := GetUserTokens(&GetUserTokensRequest{
		ID: newUser.ID,
	})

	c.Assert(err, IsNil)
	c.Assert(tokens, NotNil)
	c.Check(tokens, HasLen, 1)
	c.Check(tokens, DeepEquals, []string{token})
}
