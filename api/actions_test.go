package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestActions(t *testing.T) {
	Convey("Actions API", t, func() {
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

		project, _, err := tasks.CreateProject(&tasks.CreateProjectRequest{
			Name:        "Test Project",
			Description: "Testing",
			URL:         "https://github.com/SierraSoftworks/chieftan-server",
		})
		So(err, ShouldBeNil)
		So(project, ShouldNotBeNil)

		action, _, err := tasks.CreateAction(&tasks.CreateActionRequest{
			Name:        "Test Project",
			Description: "This is a test project",
			Variables: map[string]string{
				"x": "1",
			},
			Configurations: []models.ActionConfiguration{
				models.ActionConfiguration{
					Name: "Config 2",
					Variables: map[string]string{
						"x": "2",
					},
				},
			},
			HTTP: &models.Request{
				Method:  "GET",
				URL:     "https://github.com/SierraSoftworks/chieftan-server",
				Headers: map[string]string{},
			},
			Project: project.Summary(),
		})
		So(err, ShouldBeNil)
		So(action, ShouldNotBeNil)

		Convey("/v1/project/{project}/actions", func() {
			url := fmt.Sprintf("%s/v1/project/%s/actions", ts.URL, project.ID.Hex())

			Convey("GET", func() {
				req, err := http.NewRequest("GET", url, nil)
				So(err, ShouldBeNil)

				Convey("When not signed in", func() {
					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 401)
				})

				Convey("When signed in", func() {
					req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

					Convey("Without project/:project permissions", func() {
						res, err := http.DefaultClient.Do(req)
						defer res.Body.Close()
						So(err, ShouldBeNil)

						So(res.StatusCode, ShouldEqual, 403)
					})

					Convey("With project/:project permissions", func() {
						_, _, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{fmt.Sprintf("project/%s", project.ID.Hex())},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var as []models.Action
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&as), ShouldBeNil)

						So(as, ShouldResemble, []models.Action{*action})
					})

				})
			})

			Convey("POST", func() {
				reqBody := bytes.NewBuffer([]byte{})
				reqBodyWriter := bufio.NewWriter(reqBody)
				enc := json.NewEncoder(reqBodyWriter)
				So(enc.Encode(&tasks.CreateActionRequest{
					Name:        "Test Project",
					Description: "This is a test project",
					Variables: map[string]string{
						"x": "1",
					},
					Configurations: []models.ActionConfiguration{
						models.ActionConfiguration{
							Name: "Config 2",
							Variables: map[string]string{
								"x": "2",
							},
						},
					},
					HTTP: &models.Request{
						Method: "GET",
						URL:    "https://github.com/SierraSoftworks/chieftan-server",
					},
				}), ShouldBeNil)
				reqBodyWriter.Flush()

				req, err := http.NewRequest("POST", url, bufio.NewReader(reqBody))
				So(err, ShouldBeNil)

				Convey("When not signed in", func() {
					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 401)
				})

				Convey("When signed in", func() {
					req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

					Convey("Without project/:project/admin permissions", func() {
						res, err := http.DefaultClient.Do(req)
						defer res.Body.Close()
						So(err, ShouldBeNil)

						So(res.StatusCode, ShouldEqual, 403)
					})

					Convey("With project/:project/admin permissions", func() {
						_, _, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{fmt.Sprintf("project/%s/admin", project.ID.Hex())},
						})
						So(err, ShouldBeNil)
						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var action models.Action
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&action), ShouldBeNil)

						So(action.ID, ShouldNotBeEmpty)

					})
				})
			})
		})

		Convey("/v1/action/{action}", func() {
			url := fmt.Sprintf("%s/v1/action/%s", ts.URL, action.ID.Hex())

			Convey("GET", func() {
				req, err := http.NewRequest("GET", url, nil)
				So(err, ShouldBeNil)

				Convey("When not signed in", func() {
					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 401)
				})

				Convey("When signed in", func() {
					req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

					Convey("Without project/:project permissions", func() {
						res, err := http.DefaultClient.Do(req)
						defer res.Body.Close()
						So(err, ShouldBeNil)

						So(res.StatusCode, ShouldEqual, 403)
					})

					Convey("With project/:project permissions", func() {
						_, _, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{fmt.Sprintf("project/%s", project.ID.Hex())},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var a models.Action
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&a), ShouldBeNil)

						So(&a, ShouldResemble, action)
					})

				})
			})

			Convey("PUT", func() {
				reqBody := bytes.NewBuffer([]byte{})
				reqBodyWriter := bufio.NewWriter(reqBody)
				enc := json.NewEncoder(reqBodyWriter)
				So(enc.Encode(&tasks.CreateActionRequest{
					Variables: map[string]string{
						"x": "5",
					},
				}), ShouldBeNil)
				reqBodyWriter.Flush()

				req, err := http.NewRequest("PUT", url, bufio.NewReader(reqBody))
				So(err, ShouldBeNil)

				Convey("When not signed in", func() {
					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 401)
				})

				Convey("When signed in", func() {
					req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

					Convey("Without project/:project/admin permissions", func() {
						res, err := http.DefaultClient.Do(req)
						defer res.Body.Close()
						So(err, ShouldBeNil)

						So(res.StatusCode, ShouldEqual, 403)
					})

					Convey("With project/:project/admin permissions", func() {
						_, _, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{fmt.Sprintf("project/%s/admin", project.ID.Hex())},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var a models.Action
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&a), ShouldBeNil)

						So(a.Variables, ShouldResemble, map[string]string{
							"x": "5",
						})
					})
				})
			})
		})
	})

}
