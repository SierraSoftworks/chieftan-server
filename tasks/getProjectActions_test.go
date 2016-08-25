package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetProjectActions(t *testing.T) {
	Convey("GetProjectActions", t, func() {
		testSetup()

		newProject, _, err := CreateProject(&CreateProjectRequest{
			Name:        "Test Project",
			Description: "Test",
			URL:         "https://github.com/sierrasoftworks/chieftan-server",
		})
		So(err, ShouldBeNil)
		So(newProject, ShouldNotBeNil)

		action, _, err := CreateAction(&CreateActionRequest{
			Name:           "Test Action",
			Description:    "Test",
			Variables:      map[string]string{},
			Configurations: []models.ActionConfiguration{},
			Project:        newProject.Summary(),
		})
		So(err, ShouldBeNil)
		So(action, ShouldNotBeNil)

		actions, err := GetProjectActions(&GetProjectActionsRequest{
			ProjectID: newProject.ID.Hex(),
		})
		So(err, ShouldBeNil)
		So(actions, ShouldNotBeNil)
		So(actions, ShouldHaveLength, 1)
		So(actions[0], ShouldResemble, *action)
	})
}
