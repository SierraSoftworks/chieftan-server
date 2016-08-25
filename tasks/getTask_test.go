package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTask(t *testing.T) {
	Convey("GetTask", t, func() {
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

		newTask, _, err := CreateTask(&CreateTaskRequest{
			Metadata: &models.TaskMetadata{
				Description: "Test task",
			},
			Action:  action.Summary(),
			Project: project.Summary(),
		})
		So(err, ShouldBeNil)
		So(newTask, ShouldNotBeNil)

		task, err := GetTask(&GetTaskRequest{
			TaskID: newTask.ID.Hex(),
		})
		So(err, ShouldBeNil)
		So(task, ShouldNotBeNil)
	})
}
