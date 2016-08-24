package tasks

import (
	"testing"

	"github.com/SierraSoftworks/girder/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateToken(t *testing.T) {
	Convey("CreateToken", t, func() {
		testSetup()

		req := &CreateTokenRequest{
			UserID: "invalid user ID",
		}

		_, audit, err := CreateToken(req)
		So(err, ShouldNotBeNil)
		So(errors.From(err).Code, ShouldEqual, 400)

		user, _, err := CreateUser(&CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		req.UserID = user.ID
		token, audit, err := CreateToken(req)
		So(err, ShouldBeNil)
		So(audit, ShouldNotBeNil)
		So(audit.User, ShouldResemble, user.Summary())
		So(token, ShouldNotEqual, "")
		So(token, ShouldHaveLength, 32)
	})
}
