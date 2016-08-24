package tasks

import (
	"testing"

	"github.com/SierraSoftworks/girder/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegisterToken(t *testing.T) {
	Convey("RegisterToken", t, func() {
		testSetup()
		req := &RegisterTokenRequest{
			UserID: "invalid user ID",
		}

		_, _, err := RegisterToken(req)
		So(err, ShouldNotBeNil)
		So(errors.From(err).Code, ShouldEqual, 400)

		user, _, err := CreateUser(&CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		req.UserID = user.ID
		token, _, err := RegisterToken(req)
		So(err, ShouldNotBeNil)
		So(errors.From(err).Code, ShouldEqual, 400)

		req.Token = "0123456789abcdef0123456789abcdef"
		token, audit, err := RegisterToken(req)
		So(err, ShouldBeNil)
		So(audit, ShouldNotBeNil)
		So(audit.User, ShouldResemble, user.Summary())
		So(token, ShouldNotEqual, "")
		So(token, ShouldHaveLength, 32)
	})
}
