package tasks

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetUserTokens(t *testing.T) {
	Convey("GetUserTokens", t, func() {
		testSetup()
		newUser, _, err := CreateUser(&CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
		})

		So(err, ShouldBeNil)

		token, _, err := CreateToken(&CreateTokenRequest{
			UserID: newUser.ID,
		})

		tokens, audit, err := GetUserTokens(&GetUserTokensRequest{
			ID: newUser.ID,
		})

		So(err, ShouldBeNil)
		So(audit, ShouldNotBeNil)
		So(audit.User, ShouldNotBeNil)
		So(audit.User.Name, ShouldEqual, "Test User")
		So(audit.User.Email, ShouldEqual, "test@test.com")
		So(tokens, ShouldNotBeNil)
		So(tokens, ShouldHaveLength, 1)
		So(tokens, ShouldResemble, []string{token})
	})
}
