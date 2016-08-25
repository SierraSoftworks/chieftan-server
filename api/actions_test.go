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
	Convey("/v1/project/{project}/actions", t, func() {
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

		Convey("GET", func() {
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

			req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/project/%s/actions", ts.URL, project.ID.Hex()), nil)
			So(err, ShouldBeNil)

			Convey("When not signed in", func() {
				res, err := http.DefaultClient.Do(req)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 401)
			})

			Convey("When signed in", func() {
				req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

				res, err := http.DefaultClient.Do(req)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 200)

				var as []models.Action
				dec := json.NewDecoder(res.Body)
				So(dec.Decode(&as), ShouldBeNil)

				So(as, ShouldResemble, []models.Action{*action})
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

			req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/project/%s/actions", ts.URL, project.ID.Hex()), bufio.NewReader(reqBody))
			So(err, ShouldBeNil)

			Convey("When not signed in", func() {
				res, err := http.DefaultClient.Do(req)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 401)
			})

			Convey("When signed in", func() {
				req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

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

	Convey("/v1/action/{action}", t, func() {
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

		Convey("GET", func() {
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/action/%s", ts.URL, action.ID.Hex()), nil)
			So(err, ShouldBeNil)

			Convey("When not signed in", func() {
				res, err := http.DefaultClient.Do(req)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 401)
			})

			Convey("When signed in", func() {
				req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

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
}
