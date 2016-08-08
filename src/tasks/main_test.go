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

func (s *TasksSuite) SetUpSuite(c *C) {
	_, err := models.DB().Users().RemoveAll(&bson.M{})
	c.Assert(err, IsNil)
}
