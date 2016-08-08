package models

import . "gopkg.in/check.v1"

type AccessTokenSuite struct{}

var _ = Suite(&AccessTokenSuite{})

func (s *AccessTokenSuite) TestIsWellFormattedAccessToken(c *C) {
	c.Check(IsWellFormattedAccessToken("abc"), Equals, false)
	c.Check(IsWellFormattedAccessToken("x"), Equals, false)
	c.Check(IsWellFormattedAccessToken("0123456789abcdef0123456789abcdef"), Equals, true)
}

func (s *AccessTokenSuite) TestGenerateAccessToken(c *C) {
	token, err := GenerateAccessToken()

	c.Assert(err, IsNil)

	c.Check(token, Not(Equals), "")
	c.Check(token, HasLen, 32)

	c.Check(token, Not(Equals), "00000000000000000000000000000000")
}
