package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetTasks(t *testing.T) {
	Convey("GetTasks", t, func() {
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

		_, _, err = CreateTask(&CreateTaskRequest{
			Metadata: &models.TaskMetadata{
				Description: "Test task",
			},
			Action:  action.Summary(),
			Project: project.Summary(),
		})
		So(err, ShouldBeNil)

		_, _, err = CreateTask(&CreateTaskRequest{
			Metadata: &models.TaskMetadata{
				Description: "Test task 2",
			},
			Action:  action.Summary(),
			Project: project.Summary(),
		})
		So(err, ShouldBeNil)

		Convey("Global", func() {
			tasks, err := GetTasks(&GetTasksRequest{})
			So(err, ShouldBeNil)
			So(tasks, ShouldNotBeNil)
			So(tasks, ShouldHaveLength, 2)
		})

		Convey("Project", func() {
			tasks, err := GetTasks(&GetTasksRequest{
				ProjectID: project.ID.Hex(),
			})
			So(err, ShouldBeNil)
			So(tasks, ShouldNotBeNil)
			So(tasks, ShouldHaveLength, 2)
		})

		Convey("Action", func() {
			tasks, err := GetTasks(&GetTasksRequest{
				ActionID: action.ID.Hex(),
			})
			So(err, ShouldBeNil)
			So(tasks, ShouldNotBeNil)
			So(tasks, ShouldHaveLength, 2)
		})

		Convey("Skip", func() {
			tasks, err := GetTasks(&GetTasksRequest{
				Skip: 1,
			})
			So(err, ShouldBeNil)
			So(tasks, ShouldNotBeNil)
			So(tasks, ShouldHaveLength, 1)
		})

		Convey("Limit", func() {
			tasks, err := GetTasks(&GetTasksRequest{
				Limit: 1,
			})
			So(err, ShouldBeNil)
			So(tasks, ShouldNotBeNil)
			So(tasks, ShouldHaveLength, 1)
		})
	})
}
