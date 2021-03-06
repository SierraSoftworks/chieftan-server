package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestProjects(t *testing.T) {
	Convey("Projects API", t, func() {

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

		Convey("/v1/projects", func() {
			url := fmt.Sprintf("%s/v1/projects", ts.URL)

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

					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 200)

					Convey("When there are no projects", func() {
						_, err := models.DB().Projects().RemoveAll(&bson.M{})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var projects []models.Project
						So(json.NewDecoder(res.Body).Decode(&projects), ShouldBeNil)
						So(projects, ShouldNotBeNil)
						So(projects, ShouldHaveLength, 0)
					})
				})
			})

			Convey("POST", func() {
				reqBody := bytes.NewBuffer([]byte{})
				reqBodyWriter := bufio.NewWriter(reqBody)
				enc := json.NewEncoder(reqBodyWriter)
				So(enc.Encode(&tasks.CreateProjectRequest{
					Name:        "Test Project",
					Description: "This is a test project",
					URL:         "https://github.com/sierrasoftworks/chieftan-server",
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

					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 200)

					var project models.Project
					dec := json.NewDecoder(res.Body)
					So(dec.Decode(&project), ShouldBeNil)

					So(project.ID, ShouldNotBeEmpty)

					Convey("Grants user project permissions", func() {
						user, err := tasks.GetUser(&tasks.GetUserRequest{
							UserID: user.ID,
						})
						So(err, ShouldBeNil)
						So(user.Permissions, ShouldContain, fmt.Sprintf("project/%s", project.ID.Hex()))
						So(user.Permissions, ShouldContain, fmt.Sprintf("project/%s/admin", project.ID.Hex()))
					})
				})
			})
		})

		Convey("/v1/project/{project}", func() {
			url := fmt.Sprintf("%s/v1/project/%s", ts.URL, project.ID.Hex())

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

					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 200)
				})
			})
		})
	})
}
