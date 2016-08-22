package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestGetUser(c *C) {
	newUser, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})

	c.Assert(err, IsNil)

	user, err := GetUser(&GetUserRequest{
		ID: newUser.ID,
	})

	c.Assert(err, IsNil)
	c.Assert(user, NotNil)
	c.Check(user.ID, Equals, newUser.ID)
}
