package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func TestGetAuditLogEntry(t *testing.T) {
	Convey("GetAuditLogEntry", t, func() {
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

		entry, err := GetAuditLogEntry(&GetAuditLogEntryRequest{
			ID: newEntry.ID.Hex(),
		})

		So(err, ShouldBeNil)
		So(entry, ShouldNotBeNil)
		So(entry.ID, ShouldEqual, newEntry.ID)
		So(entry.Token, ShouldEqual, newEntry.Token)
		So(entry.User, ShouldResemble, newEntry.User)

		_, err = GetAuditLogEntry(&GetAuditLogEntryRequest{
			ID: bson.NewObjectId().Hex(),
		})
		So(err, ShouldNotBeNil)
		So(errors.From(err).Code, ShouldEqual, 404)
	})
}
