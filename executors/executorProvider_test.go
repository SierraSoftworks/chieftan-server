package executors

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestExecutorProvider(t *testing.T) {
	Convey("ExecutorProvider", t, func() {
		configuration := &models.ActionConfiguration{
			Name:      "Test",
			Variables: map[string]string{},
		}

		action := &models.Action{
			Variables: map[string]string{},
			Configurations: []models.ActionConfiguration{
				*configuration,
			},
		}

		task := &models.Task{
			Variables: map[string]string{},
			Action:    action.Summary(),
		}

		variables := map[string]string{}

		execution, err := NewExecution(&Options{
			Configuration: configuration,
			Action:        action,
			Task:          task,
			Variables:     variables,
		})
		So(err, ShouldBeNil)
		So(execution, ShouldNotBeNil)

		Convey("DefaultExecutionProvider", func() {
			ep := &DefaultExecutorProvider{}

			Convey("With No Operations", func() {
				executors := ep.GetExecutors(execution)
				So(executors, ShouldHaveLength, 0)
			})

			Convey("With an HTTP Operation", func() {
				action.HTTP = &models.Request{
					Method: "GET",
					URL:    "https://github.com/SierraSoftworks/chieftan-server",
				}
				executors := ep.GetExecutors(execution)
				So(executors, ShouldHaveLength, 1)
				So(executors[0], ShouldHaveSameTypeAs, &HTTP{})
			})
		})
	})
}
