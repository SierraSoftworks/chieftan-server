package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateTask(t *testing.T) {
	Convey("CreateTask", t, func() {
		testSetup()

		project, _, err := CreateProject(&CreateProjectRequest{
			Name:        "Test Project",
			Description: "Test",
			URL:         "https://github.com/sierrasoftworks/chieftan-server",
		})
		So(err, ShouldBeNil)
		So(project, ShouldNotBeNil)

		action, _, err := CreateAction(&CreateActionRequest{
			Name:           "Test Action",
			Description:    "Test",
			Variables:      map[string]string{},
			Configurations: []models.ActionConfiguration{},
			Project:        project.Summary(),
		})
		So(err, ShouldBeNil)
		So(action, ShouldNotBeNil)

		req := &CreateTaskRequest{
			Metadata: &models.TaskMetadata{
				Description: "Test task",
			},
			Action:  action.Summary(),
			Project: project.Summary(),
			Variables: map[string]string{
				"x": "1",
			},
		}

		task, audit, err := CreateTask(req)
		So(err, ShouldBeNil)
		So(audit, ShouldNotBeNil)
		So(audit.Task, ShouldNotBeNil)
		So(audit.Task.ID, ShouldEqual, task.ID)

		So(task, ShouldNotBeNil)
		So(task.ID, ShouldNotBeEmpty)

		Convey("Updates database", func() {
			task, err := GetTask(&GetTaskRequest{
				TaskID: task.ID.Hex(),
			})
			So(err, ShouldBeNil)
			So(task, ShouldNotBeNil)
		})
	})
}
