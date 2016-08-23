package models

import (
	"testing"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *TestSuite) SetUpTest(c *C) {
	Connect("mongodb://localhost/chieftan_test")

	_, err := DB().Users().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)

	_, err = DB().Projects().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)

	_, err = DB().Actions().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)

	_, err = DB().Tasks().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)

	_, err = DB().AuditLogs().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)
}
