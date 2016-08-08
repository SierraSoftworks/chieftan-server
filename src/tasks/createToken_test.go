package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestCreateToken(c *C) {
	req := &CreateTokenRequest{
		UserID: "invalid user ID",
	}

	_, err := CreateToken(req)
	c.Check(err, NotNil)

	switch err.(type) {
	case *TaskError:
		e := err.(*TaskError)
		c.Check(e.Code, Equals, 400)
	}

	user, err := CreateUser(&CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})
	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	req.UserID = user.ID
	token, err := CreateToken(req)
	c.Assert(err, IsNil)
	c.Check(token, Not(Equals), "")
	c.Check(token, HasLen, 32)
}
