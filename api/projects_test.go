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

func TestProjects(t *testing.T) {
	Convey("/v1/projects", t, func() {
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
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/projects", ts.URL), nil)
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

			req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/projects", ts.URL), bufio.NewReader(reqBody))
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
			})
		})
	})
}
