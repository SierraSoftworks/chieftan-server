package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestSetPermissions(c *C) {
	user, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})
	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	user, err = SetPermissions(&SetPermissionsRequest{
		UserID:      user.ID,
		Permissions: []string{"test"},
	})
	c.Assert(err, IsNil)
	c.Check(user, NotNil)

	c.Assert(user.Permissions, DeepEquals, []string{"test"})
}
