package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSetPermissions(t *testing.T) {
	Convey("SetPermissions", t, func() {
		testSetup()

		user, _, err := CreateUser(&CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		audit, err := SetPermissions(&SetPermissionsRequest{
			UserID:      user.ID,
			Permissions: []string{"test"},
		})
		So(err, ShouldBeNil)
		So(audit, ShouldNotBeNil)
		So(audit.User, ShouldResemble, user.Summary())

		user, err = GetUser(&GetUserRequest{
			ID: models.DeriveID("test@test.com"),
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		So(user.Permissions, ShouldResemble, []string{"test"})
	})
}
