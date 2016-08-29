package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAccessToken(t *testing.T) {
	Convey("AccessToken", t, func() {
		testSetup()

		Convey("IsWellFormattedAccessToken", func() {
			Convey("Too Short", func() {
				So(IsWellFormattedAccessToken("abc"), ShouldBeFalse)
				So(IsWellFormattedAccessToken("x"), ShouldBeFalse)
			})

			Convey("Invalid Character", func() {
				So(IsWellFormattedAccessToken("x123456789abcdef0123456789abcdef"), ShouldBeFalse)
			})

			Convey("Valid", func() {
				So(IsWellFormattedAccessToken("0123456789abcdef0123456789abcdef"), ShouldBeTrue)
			})
		})

		Convey("GenerateAccessToken", func() {
			token, err := GenerateAccessToken()
			So(err, ShouldBeNil)
			So(token, ShouldNotBeEmpty)
			So(token, ShouldHaveLength, 32)
			So(token, ShouldNotEqual, "00000000000000000000000000000000")
		})
	})
}
