package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestRegisterToken(c *C) {
	req := &RegisterTokenRequest{
		UserID: "invalid user ID",
	}

	_, err := RegisterToken(req)
	c.Check(err, NotNil)

	switch err.(type) {
	case *TaskError:
		e := err.(*TaskError)
		c.Check(e.Code, Equals, 400)
	default:
		c.Fail()
	}

	user, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})
	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	req.UserID = user.ID
	token, err := RegisterToken(req)
	c.Check(err, NotNil)

	switch err.(type) {
	case *TaskError:
		e := err.(*TaskError)
		c.Check(e.Code, Equals, 400)
	default:
		c.Fail()
	}

	req.Token = "0123456789abcdef0123456789abcdef"
	token, err = RegisterToken(req)
	c.Assert(err, IsNil)
	c.Check(token, Not(Equals), "")
	c.Check(token, HasLen, 32)
}
