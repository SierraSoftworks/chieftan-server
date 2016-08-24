package tasks

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetUsers(t *testing.T) {
	Convey("GetUsers", t, func() {
		testSetup()

		_, _, err := CreateUser(&CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
		})

		So(err, ShouldBeNil)

		users, err := GetUsers(&GetUsersRequest{})

		So(err, ShouldBeNil)
		So(users, ShouldNotBeNil)
		So(users, ShouldHaveLength, 1)
	})
}
