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

func TestTasks(t *testing.T) {
	Convey("Tasks API", t, func() {
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

		task, _, err := tasks.CreateTask(&tasks.CreateTaskRequest{
			Metadata: &models.TaskMetadata{
				Description: "Test task",
			},
			Variables: map[string]string{
				"x": "7",
			},
			Project: project.Summary(),
			Action:  action.Summary(),
		})
		So(err, ShouldBeNil)
		So(action, ShouldNotBeNil)

		Convey("/v1/tasks", func() {
			url := fmt.Sprintf("%s/v1/tasks", ts.URL)

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
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var items []models.Task
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&items), ShouldBeNil)

						So(items, ShouldNotBeNil)
						So(items, ShouldHaveLength, 1)
						So(items[0].Metadata, ShouldResemble, task.Metadata)
						So(items[0].Project, ShouldResemble, task.Project)
						So(items[0].Action, ShouldResemble, task.Action)
					})

				})
			})
		})

		Convey("/v1/task/{task}", func() {
			url := fmt.Sprintf("%s/v1/task/%s", ts.URL, task.ID.Hex())

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
						_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{fmt.Sprintf("project/%s", project.ID.Hex())},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var item models.Task
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&item), ShouldBeNil)

						So(item, ShouldNotBeNil)
						So(item.Metadata, ShouldResemble, task.Metadata)
						So(item.Project, ShouldResemble, task.Project)
						So(item.Action, ShouldResemble, task.Action)
					})

				})
			})
		})

		Convey("/v1/task/{task}/run", func() {
			url := fmt.Sprintf("%s/v1/task/%s/run", ts.URL, task.ID.Hex())

			Convey("POST", func() {
				reqBody := bytes.NewBuffer([]byte{})
				reqBodyWriter := bufio.NewWriter(reqBody)
				enc := json.NewEncoder(reqBodyWriter)
				So(enc.Encode(&tasks.RunTaskRequest{
					Configuration: "Config 2",
					Variables:     map[string]string{},
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

					Convey("Without project/:project permissions", func() {
						res, err := http.DefaultClient.Do(req)
						defer res.Body.Close()
						So(err, ShouldBeNil)

						So(res.StatusCode, ShouldEqual, 403)
					})

					Convey("With project/:project permissions", func() {
						_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{fmt.Sprintf("project/%s", project.ID.Hex())},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var item models.Task
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&item), ShouldBeNil)

						So(item, ShouldNotBeNil)
						So(item.Metadata, ShouldResemble, task.Metadata)
						So(item.Project, ShouldResemble, task.Project)
						So(item.Action, ShouldResemble, task.Action)
					})

				})
			})
		})

		Convey("/v1/project/{project}/tasks", func() {
			url := fmt.Sprintf("%s/v1/project/%s/tasks", ts.URL, project.ID.Hex())

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
						_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{fmt.Sprintf("project/%s", project.ID.Hex())},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var items []models.Task
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&items), ShouldBeNil)

						So(items, ShouldNotBeNil)
						So(items, ShouldHaveLength, 1)
						So(items[0].Metadata, ShouldResemble, task.Metadata)
						So(items[0].Project, ShouldResemble, task.Project)
						So(items[0].Action, ShouldResemble, task.Action)
					})

				})
			})
		})

		Convey("/v1/project/{project}/tasks/recent", func() {
			url := fmt.Sprintf("%s/v1/project/%s/tasks/recent", ts.URL, project.ID.Hex())

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
						_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{fmt.Sprintf("project/%s", project.ID.Hex())},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var items []models.Task
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&items), ShouldBeNil)

						So(items, ShouldNotBeNil)
						So(items, ShouldHaveLength, 1)
						So(items[0].Metadata, ShouldResemble, task.Metadata)
						So(items[0].Project, ShouldResemble, task.Project)
						So(items[0].Action, ShouldResemble, task.Action)
					})

				})
			})
		})

		Convey("/v1/action/{action}/tasks", func() {
			url := fmt.Sprintf("%s/v1/action/%s/tasks", ts.URL, action.ID.Hex())

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
						_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{fmt.Sprintf("project/%s", project.ID.Hex())},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var items []models.Task
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&items), ShouldBeNil)

						So(items, ShouldNotBeNil)
						So(items, ShouldHaveLength, 1)
						So(items[0].Metadata, ShouldResemble, task.Metadata)
						So(items[0].Project, ShouldResemble, task.Project)
						So(items[0].Action, ShouldResemble, task.Action)
					})

				})
			})

			Convey("POST", func() {
				reqBody := bytes.NewBuffer([]byte{})
				reqBodyWriter := bufio.NewWriter(reqBody)
				enc := json.NewEncoder(reqBodyWriter)
				So(enc.Encode(&tasks.CreateTaskRequest{
					Metadata: &models.TaskMetadata{
						Description: "Test task",
					},
					Variables: map[string]string{
						"x": "1",
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

					_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
						UserID:      user.ID,
						Permissions: []string{fmt.Sprintf("project/%s", project.ID.Hex())},
					})
					So(err, ShouldBeNil)

					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 200)

					var item models.Task
					dec := json.NewDecoder(res.Body)
					So(dec.Decode(&item), ShouldBeNil)

					So(item, ShouldNotBeNil)
					So(item.Metadata, ShouldResemble, task.Metadata)
					So(item.Project, ShouldResemble, task.Project)
					So(item.Action, ShouldResemble, task.Action)

				})
			})
		})

		Convey("/v1/action/{action}/tasks/recent", func() {
			url := fmt.Sprintf("%s/v1/action/%s/tasks/recent", ts.URL, action.ID.Hex())

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
						_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{fmt.Sprintf("project/%s", project.ID.Hex())},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var items []models.Task
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&items), ShouldBeNil)

						So(items, ShouldNotBeNil)
						So(items, ShouldHaveLength, 1)
						So(items[0].Metadata, ShouldResemble, task.Metadata)
						So(items[0].Project, ShouldResemble, task.Project)
						So(items[0].Action, ShouldResemble, task.Action)
					})

				})
			})
		})

		Convey("/v1/action/{action}/task/latest", func() {
			url := fmt.Sprintf("%s/v1/action/%s/task/latest", ts.URL, action.ID.Hex())

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
						_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{fmt.Sprintf("project/%s", project.ID.Hex())},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var item models.Task
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&item), ShouldBeNil)

						So(item, ShouldNotBeNil)
						So(item.Metadata, ShouldResemble, task.Metadata)
						So(item.Project, ShouldResemble, task.Project)
						So(item.Action, ShouldResemble, task.Action)
					})

				})
			})
		})

		Convey("/v1/action/{action}/task/latest/run", func() {
			url := fmt.Sprintf("%s/v1/action/%s/task/latest/run", ts.URL, action.ID.Hex())

			Convey("POST", func() {
				reqBody := bytes.NewBuffer([]byte{})
				reqBodyWriter := bufio.NewWriter(reqBody)
				enc := json.NewEncoder(reqBodyWriter)
				So(enc.Encode(&tasks.RunTaskRequest{
					Configuration: "Config 2",
					Variables:     map[string]string{},
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

					Convey("Without project/:project permissions", func() {
						res, err := http.DefaultClient.Do(req)
						defer res.Body.Close()
						So(err, ShouldBeNil)

						So(res.StatusCode, ShouldEqual, 403)
					})

					Convey("With project/:project permissions", func() {
						_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{fmt.Sprintf("project/%s", project.ID.Hex())},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						So(err, ShouldBeNil)
						So(res.StatusCode, ShouldEqual, 200)

						var item models.Task
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&item), ShouldBeNil)

						So(item, ShouldNotBeNil)
						So(item.Metadata, ShouldResemble, task.Metadata)
						So(item.Project, ShouldResemble, task.Project)
						So(item.Action, ShouldResemble, task.Action)
					})

				})
			})
		})
	})
}
