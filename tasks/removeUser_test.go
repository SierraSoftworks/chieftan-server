package tasks

import (
	"testing"

	"github.com/SierraSoftworks/girder/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveUser(t *testing.T) {
	Convey("RemoveUser", t, func() {
		testSetup()
		user, _, err := CreateUser(&CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		removedUser, audit, err := RemoveUser(&RemoveUserRequest{
			UserID: user.ID,
		})
		So(err, ShouldBeNil)
		So(removedUser, ShouldNotBeNil)
		So(removedUser, ShouldResemble, user)
		So(audit, ShouldNotBeNil)
		So(audit.User, ShouldResemble, user.Summary())

		Convey("Updates database", func() {
			user, err := GetUser(&GetUserRequest{
				UserID: user.ID,
			})
			So(err, ShouldNotBeNil)
			So(user, ShouldBeNil)
			So(errors.From(err).Code, ShouldEqual, 404)
		})
	})
}
