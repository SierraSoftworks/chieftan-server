package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	. "gopkg.in/check.v1"
)

func (s *TasksSuite) TestGetAuditLog(c *C) {
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

	entries, err := GetAuditLog(&GetAuditLogRequest{})

	c.Assert(err, IsNil)
	c.Assert(entries, NotNil)
	c.Assert(entries, HasLen, 1)

	entry := entries[0]
	c.Check(entry.ID, Equals, newEntry.ID)
	c.Check(entry.Token, Equals, newEntry.Token)
	c.Check(entry.User, DeepEquals, newEntry.User)
}
