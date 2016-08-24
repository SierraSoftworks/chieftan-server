package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SierraSoftworks/chieftan-server/tasks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTokens(t *testing.T) {
	Convey("/v1/tokens", t, func() {
		setUpTest()
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

		Convey("DELETE", func() {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1/tokens", ts.URL), nil)
			So(err, ShouldBeNil)

			Convey("When not signed in", func() {
				res, err := http.DefaultClient.Do(req)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 401)
			})

			Convey("When signed in", func() {

				req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

				Convey("Without admin permissions", func() {
					res, err := http.DefaultClient.Do(req)
					defer res.Body.Close()
					So(err, ShouldBeNil)

					So(res.StatusCode, ShouldEqual, 403)
				})

				Convey("With admin permissions", func() {
					_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
						UserID:      user.ID,
						Permissions: []string{"admin"},
					})
					So(err, ShouldBeNil)

					res, err := http.DefaultClient.Do(req)
					defer res.Body.Close()
					So(err, ShouldBeNil)

					So(res.StatusCode, ShouldEqual, 200)
				})
			})
		})
	})

	Convey("/v1/token/{token}", t, func() {
		setUpTest()
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

		Convey("DELETE", func() {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1/token/%s", ts.URL, token), nil)
			So(err, ShouldBeNil)

			Convey("When not signed in", func() {
				res, err := http.DefaultClient.Do(req)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 401)
			})

			Convey("When signed in", func() {

				req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

				Convey("Without admin/users permissions", func() {
					res, err := http.DefaultClient.Do(req)
					defer res.Body.Close()
					So(err, ShouldBeNil)

					So(res.StatusCode, ShouldEqual, 403)
				})

				Convey("With admin/users permissions", func() {
					_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
						UserID:      user.ID,
						Permissions: []string{"admin/users"},
					})
					So(err, ShouldBeNil)

					res, err := http.DefaultClient.Do(req)
					defer res.Body.Close()
					So(err, ShouldBeNil)

					So(res.StatusCode, ShouldEqual, 200)
				})
			})
		})
	})

	Convey("/v1/user/{user}/tokens", t, func() {
		setUpTest()
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

		Convey("GET", func() {
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/user/%s/tokens", ts.URL, user.ID), nil)
			So(err, ShouldBeNil)

			Convey("When not signed in", func() {
				res, err := http.DefaultClient.Do(req)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 401)
			})

			Convey("When signed in", func() {

				req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

				Convey("Without admin/users permissions", func() {
					res, err := http.DefaultClient.Do(req)
					defer res.Body.Close()
					So(err, ShouldBeNil)

					So(res.StatusCode, ShouldEqual, 403)
				})

				Convey("With admin/users permissions", func() {
					_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
						UserID:      user.ID,
						Permissions: []string{"admin/users"},
					})
					So(err, ShouldBeNil)

					res, err := http.DefaultClient.Do(req)
					defer res.Body.Close()
					So(err, ShouldBeNil)

					So(res.StatusCode, ShouldEqual, 200)

					var tokens []string
					dec := json.NewDecoder(res.Body)
					So(dec.Decode(&tokens), ShouldBeNil)
					So(tokens, ShouldHaveLength, 1)
				})
			})
		})

		Convey("POST", func() {
			req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/user/%s/tokens", ts.URL, user.ID), nil)
			So(err, ShouldBeNil)

			Convey("When not signed in", func() {
				res, err := http.DefaultClient.Do(req)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 401)
			})

			Convey("When signed in", func() {

				req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

				Convey("Without admin/users permissions", func() {
					res, err := http.DefaultClient.Do(req)
					defer res.Body.Close()
					So(err, ShouldBeNil)

					So(res.StatusCode, ShouldEqual, 403)
				})

				Convey("With admin/users permissions", func() {
					_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
						UserID:      user.ID,
						Permissions: []string{"admin/users"},
					})
					So(err, ShouldBeNil)

					res, err := http.DefaultClient.Do(req)
					defer res.Body.Close()
					So(err, ShouldBeNil)

					So(res.StatusCode, ShouldEqual, 200)

					var data struct {
						Token string `json:"token"`
					}
					dec := json.NewDecoder(res.Body)
					So(dec.Decode(&data), ShouldBeNil)
					So(data.Token, ShouldNotBeEmpty)
				})
			})
		})

		Convey("DELETE", func() {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1/user/%s/tokens", ts.URL, user.ID), nil)
			So(err, ShouldBeNil)

			Convey("When not signed in", func() {
				res, err := http.DefaultClient.Do(req)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 401)
			})

			Convey("When signed in", func() {

				req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

				Convey("Without admin/users permissions", func() {
					res, err := http.DefaultClient.Do(req)
					defer res.Body.Close()
					So(err, ShouldBeNil)

					So(res.StatusCode, ShouldEqual, 403)
				})

				Convey("With admin/users permissions", func() {
					_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
						UserID:      user.ID,
						Permissions: []string{"admin/users"},
					})
					So(err, ShouldBeNil)

					res, err := http.DefaultClient.Do(req)
					defer res.Body.Close()
					So(err, ShouldBeNil)

					So(res.StatusCode, ShouldEqual, 200)
				})
			})
		})
	})
}
