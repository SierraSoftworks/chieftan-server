package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAddPermissions(t *testing.T) {
	Convey("AddPermissions", t, func() {
		testSetup()

		user, _, err := CreateUser(&CreateUserRequest{
			Name:        "Test User",
			Email:       "test@test.com",
			Permissions: []string{"base"},
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		user, audit, err := AddPermissions(&AddPermissionsRequest{
			UserID:      user.ID,
			Permissions: []string{"test"},
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.Permissions, ShouldResemble, []string{"base", "test"})
		So(audit, ShouldNotBeNil)
		So(audit.User, ShouldResemble, user.Summary())

		Convey("Updates database", func() {
			user, err := GetUser(&GetUserRequest{
				ID: models.DeriveID("test@test.com"),
			})
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)

			So(user.Permissions, ShouldResemble, []string{"base", "test"})
		})
	})
}
