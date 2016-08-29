package executors

import (
	"bytes"
	"strings"
	"testing"

	"net/http/httptest"

	"net/http"

	"io/ioutil"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

type testHTTPCallContext struct {
	Request *http.Request
	Data    string
}

func TestHTTP(t *testing.T) {
	Convey("HTTP", t, func() {
		configuration := &models.ActionConfiguration{
			Name:      "Test",
			Variables: map[string]string{},
		}

		action := &models.Action{
			Variables: map[string]string{
				"action": "test",
			},
			Configurations: []models.ActionConfiguration{
				*configuration,
			},
		}

		task := &models.Task{
			Variables: map[string]string{},
			Action:    action.Summary(),
		}

		variables := map[string]string{
			"TOKEN": "abcd",
		}

		execution, err := NewExecution(&Options{
			Configuration: configuration,
			Action:        action,
			Task:          task,
			Variables:     variables,
		})
		So(err, ShouldBeNil)
		So(execution, ShouldNotBeNil)

		executor := &HTTP{}

		Convey("Name", func() {
			So(executor.Name(), ShouldEqual, "HTTP")
		})

		Convey("Run", func() {
			var calledContext *testHTTPCallContext

			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				data, err := ioutil.ReadAll(r.Body)
				if err != nil {
					w.WriteHeader(500)
					return
				}

				calledContext = &testHTTPCallContext{
					Request: r,
					Data:    bytes.NewBuffer(data).String(),
				}
			})

			ts := httptest.NewServer(mux)

			Convey("GET", func() {
				action.HTTP = &models.Request{
					Method: "GET",
					URL:    ts.URL,
					Headers: map[string]string{
						"X-Token": "{{TOKEN}}",
					},
				}

				err := executor.Run(execution)
				So(err, ShouldBeNil)
				So(calledContext, ShouldNotBeNil)
				So(calledContext.Request, ShouldNotBeNil)
				So(calledContext.Request.Method, ShouldEqual, "GET")
				So(calledContext.Request.Header.Get("X-Token"), ShouldEqual, "abcd")
			})

			Convey("POST", func() {
				action.HTTP = &models.Request{
					Method: "POST",
					URL:    ts.URL,
					Headers: map[string]string{
						"X-Token": "{{TOKEN}}",
					},
				}

				Convey("With a String Payload", func() {
					action.HTTP.Data = "action={{action}}"

					err := executor.Run(execution)
					So(err, ShouldBeNil)
					So(calledContext, ShouldNotBeNil)
					So(calledContext.Request, ShouldNotBeNil)
					So(calledContext.Request.Method, ShouldEqual, "POST")
					So(calledContext.Request.Header.Get("X-Token"), ShouldEqual, "abcd")
					So(calledContext.Data, ShouldEqual, "action=test")
				})

				Convey("With a Structured Payload", func() {
					action.HTTP.Data = struct {
						Action string `json:"action"`
					}{Action: "{{action}}"}

					err := executor.Run(execution)
					So(err, ShouldBeNil)
					So(calledContext, ShouldNotBeNil)
					So(calledContext.Request, ShouldNotBeNil)
					So(calledContext.Request.Method, ShouldEqual, "POST")
					So(calledContext.Request.Header.Get("X-Token"), ShouldEqual, "abcd")
					So(strings.TrimSpace(calledContext.Data), ShouldEqual, `{"action":"test"}`)

				})
			})
		})
	})
}
