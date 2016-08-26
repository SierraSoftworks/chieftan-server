package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateAuditLogEntry(t *testing.T) {
	Convey("CreateAuditLogEntry", t, func() {
		testSetup()

		req := &CreateAuditLogEntryRequest{
			Type: "test",
			User: &models.UserSummary{
				ID:    "test",
				Name:  "Test User",
				Email: "test@test.com",
			},
			Token:   "0123456789abcdef0123456789abcdef",
			Context: &models.AuditLogContext{},
		}

		entry, err := CreateAuditLogEntry(req)
		So(err, ShouldBeNil)
		So(entry, ShouldNotBeNil)

		So(entry.ID, ShouldNotEqual, "")
		So(entry.Type, ShouldEqual, "test")
		So(entry.User, ShouldResemble, models.UserSummary{
			ID:    "test",
			Name:  "Test User",
			Email: "test@test.com",
		})
		So(entry.Token, ShouldEqual, "0123456789abcdef0123456789abcdef")
		So(entry.Context, ShouldResemble, models.AuditLogContext{})

		Convey("Updates database", func() {
			entry, err := GetAuditLogEntry(&GetAuditLogEntryRequest{
				ID: entry.ID.Hex(),
			})
			So(err, ShouldBeNil)
			So(entry, ShouldNotBeNil)
			So(entry.Token, ShouldEqual, "0123456789abcdef0123456789abcdef")
		})
	})
}
