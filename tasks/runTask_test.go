package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRunTask(t *testing.T) {
	Convey("RunTask", t, func() {
		testSetup()

		project, _, err := CreateProject(&CreateProjectRequest{
			Name:        "Test Project",
			Description: "Test",
			URL:         "https://github.com/sierrasoftworks/chieftan-server",
		})
		So(err, ShouldBeNil)
		So(project, ShouldNotBeNil)

		action, _, err := CreateAction(&CreateActionRequest{
			Name:        "Test Action",
			Description: "Test",
			Variables:   map[string]string{},
			Configurations: []models.ActionConfiguration{
				models.ActionConfiguration{
					Name:      "Test",
					Variables: map[string]string{},
				},
			},
			Project: project.Summary(),
		})
		So(err, ShouldBeNil)
		So(action, ShouldNotBeNil)

		task, _, err := CreateTask(&CreateTaskRequest{
			Metadata: &models.TaskMetadata{
				Description: "Test task",
			},
			Action:  action.Summary(),
			Project: project.Summary(),
		})
		So(err, ShouldBeNil)
		So(task, ShouldNotBeNil)

		execution, audit, err := RunTask(&RunTaskRequest{
			TaskID:        task.ID.Hex(),
			Configuration: "Test",
			Variables: map[string]string{
				"x": "1",
			},
		})

		So(err, ShouldBeNil)
		So(execution, ShouldNotBeNil)
		So(execution.Configuration, ShouldResemble, &models.ActionConfiguration{
			Name:      "Test",
			Variables: map[string]string{},
		})
		So(execution.Action, ShouldResemble, action)
		So(execution.StateChanged, ShouldNotBeNil)

		So(audit, ShouldNotBeNil)
		So(audit.Action, ShouldResemble, action.Summary())
		So(audit.Task, ShouldResemble, task.Summary())
		So(audit.Project, ShouldResemble, task.Project)

		for task := range execution.StateChanged {
			So(task, ShouldNotBeNil)
			So(task.ID.Hex(), ShouldEqual, execution.Task.ID.Hex())
		}

		So(execution.Task.State, ShouldEqual, models.TaskStatePassed)
	})
}
