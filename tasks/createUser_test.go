package tasks

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateUser(t *testing.T) {
	Convey("CreateUser", t, func() {
		testSetup()

		req := &CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
			Permissions: []string{
				"admin",
			},
		}

		user, audit, err := CreateUser(req)
		So(err, ShouldBeNil)
		So(audit, ShouldNotBeNil)
		So(audit.User, ShouldNotBeNil)
		So(audit.User.ID, ShouldEqual, user.ID)

		So(user, ShouldNotBeNil)
		So(user.ID, ShouldEqual, "b642b4217b34b1e8d3bd915fc65c4452")
		So(user.Name, ShouldEqual, "Test User")
		So(user.Email, ShouldEqual, "test@test.com")
		So(user.Permissions, ShouldResemble, []string{"admin"})
		So(user.Tokens, ShouldResemble, []string{})

		Convey("Updates database", func() {
			user, err := GetUser(&GetUserRequest{
				ID: user.ID,
			})
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)
		})
	})
}
