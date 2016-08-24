package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (s *TasksSuite) TestGetAuditLogEntry(c *C) {
	newEntry, err := CreateAuditLogEntry(&CreateAuditLogEntryRequest{
		Type: "test",
		User: &models.UserSummary{
			ID:    "test",
			Name:  "Test User",
			Email: "test@test.com",
		},
		Token:   "0123456789abcdef0123456789abcdef",
		Context: &models.AuditLogContext{},
	})

	c.Assert(err, IsNil)
	c.Assert(newEntry, NotNil)

	entry, err := GetAuditLogEntry(&GetAuditLogEntryRequest{
		ID: newEntry.ID.Hex(),
	})

	c.Assert(err, IsNil)
	c.Assert(entry, NotNil)
	c.Check(entry.ID, Equals, newEntry.ID)
	c.Check(entry.Token, Equals, newEntry.Token)
	c.Check(entry.User, DeepEquals, newEntry.User)

	_, err = GetAuditLogEntry(&GetAuditLogEntryRequest{
		ID: bson.NewObjectId().Hex(),
	})
	c.Assert(err, NotNil)
	c.Check(errors.From(err).Code, Equals, 404)
}
