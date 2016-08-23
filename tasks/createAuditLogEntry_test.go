package tasks

import . "gopkg.in/check.v1"
import "github.com/SierraSoftworks/chieftan-server/models"

func (s *TasksSuite) TestCreateAuditLogEntry(c *C) {
	req := &CreateAuditLogEntryRequest{
		Type: "test",
		User: models.UserSummary{
			ID:    "test",
			Name:  "Test User",
			Email: "test@test.com",
		},
		Token:   "0123456789abcdef0123456789abcdef",
		Context: models.AuditLogContext{},
	}

	entry, err := CreateAuditLogEntry(req)
	c.Assert(err, IsNil)
	c.Assert(entry, NotNil)

	c.Check(entry.ID, Not(Equals), "")
	c.Check(entry.Type, Equals, "test")
}
