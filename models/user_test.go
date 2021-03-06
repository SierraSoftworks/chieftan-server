package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUser(t *testing.T) {
	Convey("User", t, func() {
		Convey("UpdateID", func() {
			user := User{
				Name:  "Benjamin Pannell",
				Email: "bpannell@emss.co.za",
			}

			user.UpdateID()

			So(user.ID, ShouldEqual, "c2d8df67421f13020b46dd5bdf18b36c")
		})

		Convey("Girder API", func() {
			user := User{
				ID: "c2d8df67421f13020b46dd5bdf18b36c",
				Permissions: []string{
					"test",
				},
			}

			Convey("GetID", func() {
				So(user.GetID(), ShouldEqual, "c2d8df67421f13020b46dd5bdf18b36c")
			})

			Convey("GetPermissions", func() {
				So(user.GetPermissions(), ShouldResemble, []string{"test"})
			})
		})

		Convey("Summary", func() {
			user := User{
				ID:    "c2d8df67421f13020b46dd5bdf18b36c",
				Name:  "Benjamin Pannell",
				Email: "bpannell@emss.co.za",
				Permissions: []string{
					"admin",
					"admin/users",
				},
				Tokens: []string{
					"abcdef",
				},
			}

			summary := user.Summary()

			So(summary, ShouldNotBeNil)
			So(summary, ShouldResemble, &UserSummary{
				ID:    "c2d8df67421f13020b46dd5bdf18b36c",
				Name:  "Benjamin Pannell",
				Email: "bpannell@emss.co.za",
			})
		})

		Convey("NewToken", func() {
			user := User{}

			token, err := user.NewToken()
			So(err, ShouldBeNil)
			So(token, ShouldNotBeEmpty)
			So(user.Tokens, ShouldResemble, []string{token})
		})

		Convey("IsValidUserID", func() {
			Convey("Too Short", func() {
				So(IsValidUserID("abc"), ShouldBeFalse)
				So(IsValidUserID("x"), ShouldBeFalse)
			})

			Convey("Invalid Character", func() {
				So(IsValidUserID("x123456789abcdef0123456789abcdef"), ShouldBeFalse)
			})

			Convey("Valid", func() {
				So(IsValidUserID("0123456789abcdef0123456789abcdef"), ShouldBeTrue)
			})
		})
	})
}
