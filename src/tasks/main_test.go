package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/src/models"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func Test(t *testing.T) { TestingT(t) }

type TasksSuite struct{}

var _ = Suite(&TasksSuite{})

func (s *TasksSuite) SetUpTest(c *C) {
	_, err := models.DB().Users().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)

	_, err = models.DB().Projects().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)

	_, err = models.DB().Actions().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)

	_, err = models.DB().Tasks().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)

	_, err = models.DB().AuditLogs().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)
}
