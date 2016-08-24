package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAuditLog(t *testing.T) {
	Convey("GetAuditLog", t, func() {
		testSetup()

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

		So(err, ShouldBeNil)
		So(newEntry, ShouldNotBeNil)

		entries, err := GetAuditLog(&GetAuditLogRequest{})

		So(err, ShouldBeNil)
		So(entries, ShouldNotBeNil)
		So(entries, ShouldHaveLength, 1)

		entry := entries[0]
		So(entry.ID, ShouldEqual, newEntry.ID)
		So(entry.Token, ShouldEqual, newEntry.Token)
		So(entry.User, ShouldResemble, newEntry.User)
	})
}
