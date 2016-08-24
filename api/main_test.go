package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	"github.com/SierraSoftworks/girder"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func setUpTest() {
	models.Connect("mongodb://localhost/chieftan_test")

	_, err := models.DB().Users().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)

	_, err = models.DB().Projects().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)

	_, err = models.DB().AuditLogs().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)
}

func TestAuthentication(t *testing.T) {
	Convey("Authentication", t, func() {
		setUpTest()

		Router().Path("/v1/test/auth").Methods("GET").Handler(girder.NewHandler(func(c *girder.Context) (interface{}, error) {
			return struct{}{}, nil
		}).RequireAuthentication(getUser).RequirePermission("test"))

		ts := httptest.NewServer(Router())
		defer ts.Close()

		user, _, err := tasks.CreateUser(&tasks.CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		token, _, err := tasks.CreateToken(&tasks.CreateTokenRequest{
			UserID: user.ID,
		})
		So(err, ShouldBeNil)

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/test/auth", ts.URL), nil)
		So(err, ShouldBeNil)

		Convey("Without an auth token", func() {
			res, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, 401)
		})

		Convey("With an invalid auth token type", func() {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

			res, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, 401)
		})

		Convey("With a malformed auth token", func() {
			req.Header.Set("Authorization", "Token abc")

			res, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, 401)
		})

		Convey("With an invalid auth token", func() {
			req.Header.Set("Authorization", "Token 00000000000000000000000000000000")

			res, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, 401)
		})

		Convey("With a valid auth token", func() {
			req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

			Convey("Without permission", func() {
				res, err := http.DefaultClient.Do(req)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 403)
			})

			Convey("With permission", func() {
				_, err := tasks.SetPermissions(&tasks.SetPermissionsRequest{
					UserID:      user.ID,
					Permissions: []string{"test"},
				})
				So(err, ShouldBeNil)

				res, err := http.DefaultClient.Do(req)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 200)
			})
		})
	})
}
