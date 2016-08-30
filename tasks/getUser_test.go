package tasks

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetUser(t *testing.T) {
	Convey("GetUser", t, func() {
		testSetup()

		newUser, _, err := CreateUser(&CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
		})

		So(err, ShouldBeNil)

		user, err := GetUser(&GetUserRequest{
			UserID: newUser.ID,
		})

		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)
		So(user.ID, ShouldEqual, newUser.ID)
	})
}
