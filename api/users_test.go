package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUsers(t *testing.T) {
	Convey("Users API", t, func() {
		setUpTest()
		ts := httptest.NewServer(Router())
		defer ts.Close()

		user, _, err := tasks.CreateUser(&tasks.CreateUserRequest{
			Name:  "Test User",
			Email: "test@test.com",
		})
		So(err, ShouldBeNil)
		So(user, ShouldNotBeNil)

		Convey("/v1/users", func() {
			url := fmt.Sprintf("%s/v1/users", ts.URL)

			Convey("GET", func() {
				req, err := http.NewRequest("GET", url, nil)
				So(err, ShouldBeNil)

				Convey("When not signed in", func() {
					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 401)
				})

				Convey("When signed in", func() {
					token, _, err := tasks.CreateToken(&tasks.CreateTokenRequest{
						UserID: user.ID,
					})
					So(err, ShouldBeNil)

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

						var users []models.User
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&users), ShouldBeNil)
						So(users, ShouldHaveLength, 1)
					})
				})
			})

			Convey("POST", func() {
				reqBodyBuf := bytes.NewBuffer([]byte{})
				reqBody := bufio.NewWriter(reqBodyBuf)
				enc := json.NewEncoder(reqBody)
				So(enc.Encode(tasks.CreateUserRequest{
					Email:       "newuser@test.com",
					Name:        "Test User",
					Permissions: []string{"test"},
				}), ShouldBeNil)
				So(reqBody.Flush(), ShouldBeNil)

				Convey("When not signed in", func() {
					req, err := http.NewRequest("POST", url, bufio.NewReader(reqBodyBuf))
					So(err, ShouldBeNil)

					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 401)
				})

				Convey("When signed in", func() {
					token, _, err := tasks.CreateToken(&tasks.CreateTokenRequest{
						UserID: user.ID,
					})
					So(err, ShouldBeNil)

					Convey("Without admin/users permissions", func() {
						req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/users", ts.URL), bufio.NewReader(reqBodyBuf))
						So(err, ShouldBeNil)

						req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

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

						req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/users", ts.URL), bufio.NewReader(reqBodyBuf))
						So(err, ShouldBeNil)

						req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

						res, err := http.DefaultClient.Do(req)
						defer res.Body.Close()
						So(err, ShouldBeNil)

						So(res.StatusCode, ShouldEqual, 200)

						var u models.User
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&u), ShouldBeNil)
						So(u.Email, ShouldEqual, "newuser@test.com")
					})
				})
			})
		})

		Convey("/v1/user/{user}", func() {
			url := fmt.Sprintf("%s/v1/user/%s", ts.URL, user.ID)

			Convey("GET", func() {
				req, err := http.NewRequest("GET", url, nil)
				So(err, ShouldBeNil)

				Convey("When not signed in", func() {
					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 401)
				})

				Convey("When signed in", func() {
					token, _, err := tasks.CreateToken(&tasks.CreateTokenRequest{
						UserID: user.ID,
					})
					So(err, ShouldBeNil)

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

						var u models.User
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&u), ShouldBeNil)
						So(u.ID, ShouldEqual, user.ID)
					})
				})
			})
		})

		Convey("/v1/user", func() {
			url := fmt.Sprintf("%s/v1/user", ts.URL)

			Convey("GET", func() {
				req, err := http.NewRequest("GET", url, nil)
				So(err, ShouldBeNil)

				Convey("When not signed in", func() {
					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 401)
				})

				Convey("When signed in", func() {
					token, _, err := tasks.CreateToken(&tasks.CreateTokenRequest{
						UserID: user.ID,
					})
					So(err, ShouldBeNil)

					req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

					res, err := http.DefaultClient.Do(req)
					defer res.Body.Close()
					So(err, ShouldBeNil)

					So(res.StatusCode, ShouldEqual, 200)

					var u models.User
					dec := json.NewDecoder(res.Body)
					So(dec.Decode(&u), ShouldBeNil)
					So(u.ID, ShouldEqual, user.ID)
				})
			})
		})
	})
}
