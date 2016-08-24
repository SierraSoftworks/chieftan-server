package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	. "gopkg.in/check.v1"
)

func (s *TasksSuite) TestSetPermissions(c *C) {
	user, _, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})
	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	audit, err := SetPermissions(&SetPermissionsRequest{
		UserID:      user.ID,
		Permissions: []string{"test"},
	})
	c.Assert(err, IsNil)
	c.Assert(audit, NotNil)
	c.Check(audit.User, DeepEquals, user.Summary())

	user, err = GetUser(&GetUserRequest{
		ID: models.DeriveID("test@test.com"),
	})
	c.Assert(err, IsNil)
	c.Check(user, NotNil)

	c.Assert(user.Permissions, DeepEquals, []string{"test"})
}
