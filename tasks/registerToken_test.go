package tasks

import (
	"github.com/SierraSoftworks/girder/errors"
	. "gopkg.in/check.v1"
)

func (s *TasksSuite) TestRegisterToken(c *C) {
	req := &RegisterTokenRequest{
		UserID: "invalid user ID",
	}

	_, _, err := RegisterToken(req)
	c.Check(err, NotNil)
	c.Check(errors.From(err).Code, Equals, 400)

	user, _, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})
	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	req.UserID = user.ID
	token, _, err := RegisterToken(req)
	c.Check(err, NotNil)
	c.Check(errors.From(err).Code, Equals, 400)

	req.Token = "0123456789abcdef0123456789abcdef"
	token, audit, err := RegisterToken(req)
	c.Assert(err, IsNil)
	c.Assert(audit, NotNil)
	c.Check(token, Not(Equals), "")
	c.Check(token, HasLen, 32)
}
