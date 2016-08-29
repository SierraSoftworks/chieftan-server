package tasks

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveToken(t *testing.T) {
	Convey("RemoveToken", t, func() {
		testSetup()
		user, _, err := CreateUser(&CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		token, _, err := CreateToken(&CreateTokenRequest{
			UserID: user.ID,
		})
		So(err, ShouldBeNil)
		So(token, ShouldNotEqual, "")

		audit, err := RemoveToken(&RemoveTokenRequest{
			Token: token,
		})
		So(err, ShouldBeNil)
		So(audit, ShouldNotBeNil)
		So(audit.User, ShouldResemble, user.Summary())

		Convey("Updates database", func() {
			tokens, _, err := GetTokens(&GetTokensRequest{
				UserID: user.ID,
			})
			So(err, ShouldBeNil)
			So(tokens, ShouldNotBeNil)
			So(tokens, ShouldResemble, []string{})
		})
	})
}
