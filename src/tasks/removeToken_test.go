package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestRemoveToken(c *C) {
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

	err = RemoveToken(&RemoveTokenRequest{
		Token: token,
	})
	c.Assert(err, IsNil)
}
