package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRemovePermissions(t *testing.T) {
	Convey("RemovePermissions", t, func() {
		testSetup()

		user, _, err := CreateUser(&CreateUserRequest{
			Name:        "Test User",
			Email:       "test@test.com",
			Permissions: []string{"base"},
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		user, audit, err := RemovePermissions(&RemovePermissionsRequest{
			UserID:      user.ID,
			Permissions: []string{"base"},
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.Permissions, ShouldResemble, []string{})
		So(audit, ShouldNotBeNil)
		So(audit.User, ShouldResemble, user.Summary())

		Convey("Updates database", func() {
			user, err := GetUser(&GetUserRequest{
				UserID: models.DeriveID("test@test.com"),
			})
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)

			So(user.Permissions, ShouldResemble, []string{})
		})
	})
}
