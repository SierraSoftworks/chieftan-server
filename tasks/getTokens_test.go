package tasks

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTokens(t *testing.T) {
	Convey("GetTokens", t, func() {
		testSetup()

		user, _, err := CreateUser(&CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
		})

		So(err, ShouldBeNil)

		token, _, err := CreateToken(&CreateTokenRequest{
			UserID: user.ID,
		})

		So(err, ShouldBeNil)

		tokens, audit, err := GetTokens(&GetTokensRequest{
			UserID: user.ID,
		})

		So(err, ShouldBeNil)
		So(audit, ShouldNotBeNil)
		So(audit.User, ShouldResemble, user.Summary())
		So(tokens, ShouldNotBeNil)
		So(tokens, ShouldHaveLength, 1)
		So(tokens, ShouldResemble, []string{token})
	})
}
