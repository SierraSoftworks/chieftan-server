package api

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) SetUpTest(c *C) {
	models.Connect("mongodb://localhost/chieftan_test")

	_, err := models.DB().Users().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)

	_, err = models.DB().AuditLogs().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)
}
