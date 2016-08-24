package tasks

import (
	"github.com/SierraSoftworks/girder/errors"
	. "gopkg.in/check.v1"
)

func (s *TasksSuite) TestCreateToken(c *C) {
	req := &CreateTokenRequest{
		UserID: "invalid user ID",
	}

	_, audit, err := CreateToken(req)
	c.Assert(err, NotNil)
	c.Check(errors.From(err).Code, Equals, 400)

	user, _, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})
	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	req.UserID = user.ID
	token, audit, err := CreateToken(req)
	c.Assert(err, IsNil)
	c.Assert(audit, NotNil)
	c.Check(audit.User, DeepEquals, user.Summary())
	c.Check(token, Not(Equals), "")
	c.Check(token, HasLen, 32)
}
