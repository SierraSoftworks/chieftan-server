package models

import . "gopkg.in/check.v1"

func (s *TestSuite) TestUserUpdateID(c *C) {
	user := User{
		Name:  "Benjamin Pannell",
		Email: "bpannell@emss.co.za",
	}

	user.UpdateID()

	c.Check(user.ID, Equals, "c2d8df67421f13020b46dd5bdf18b36c")
}

func (s *TestSuite) TestUserSummary(c *C) {
	user := User{
		ID:    "c2d8df67421f13020b46dd5bdf18b36c",
		Name:  "Benjamin Pannell",
		Email: "bpannell@emss.co.za",
		Permissions: []string{
			"admin",
			"admin/users",
		},
		Tokens: []string{
			"abcdef",
		},
	}

	summary := user.Summary()
	c.Check(summary, DeepEquals, &UserSummary{
		ID:    "c2d8df67421f13020b46dd5bdf18b36c",
		Name:  "Benjamin Pannell",
		Email: "bpannell@emss.co.za",
	})
}

func (s *TestSuite) TestUserNewToken(c *C) {
	user := User{}

	token, err := user.NewToken()
	c.Assert(err, IsNil)
	c.Check(token, Not(Equals), "")
	c.Check(user.Tokens, DeepEquals, []string{token})
}

func (s *TestSuite) TestUserIsValidUserID(c *C) {
	c.Check(IsValidUserID("abc"), Equals, false)
	c.Check(IsValidUserID("x"), Equals, false)
	c.Check(IsValidUserID("0123456789abcdef0123456789abcdef"), Equals, true)
}
