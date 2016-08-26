package tasks

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveAllTokens(t *testing.T) {
	Convey("RemoveAllTokens", t, func() {
		testSetup()

		user, _, err := CreateUser(&CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		Convey("Globally", func() {
			token, _, err := CreateToken(&CreateTokenRequest{
				UserID: user.ID,
			})
			So(err, ShouldBeNil)
			So(token, ShouldNotEqual, "")

			audit, err := RemoveAllTokens(&RemoveAllTokensRequest{})
			So(err, ShouldBeNil)
			So(audit, ShouldNotBeNil)

			Convey("Updates database", func() {
				tokens, _, err := GetTokens(&GetTokensRequest{
					UserID: user.ID,
				})
				So(err, ShouldBeNil)
				So(tokens, ShouldResemble, []string{})
			})
		})

		Convey("For User", func() {
			token, _, err := CreateToken(&CreateTokenRequest{
				UserID: user.ID,
			})
			So(err, ShouldBeNil)
			So(token, ShouldNotEqual, "")

			audit, err := RemoveAllTokens(&RemoveAllTokensRequest{
				UserID: user.ID,
			})
			So(err, ShouldBeNil)
			So(audit, ShouldNotBeNil)
			So(audit.User, ShouldResemble, user.Summary())

			Convey("Updates database", func() {
				tokens, _, err := GetTokens(&GetTokensRequest{
					UserID: user.ID,
				})
				So(err, ShouldBeNil)
				So(tokens, ShouldResemble, []string{})
			})
		})
	})
}
