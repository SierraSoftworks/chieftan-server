package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestGetUsers(c *C) {
	_, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})

	c.Assert(err, IsNil)

	users, err := GetUsers(&GetUsersRequest{})

	c.Assert(err, IsNil)
	c.Assert(users, NotNil)
	c.Check(users, HasLen, 1)
}
