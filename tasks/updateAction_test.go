package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpdateAction(t *testing.T) {
	Convey("UpdateAction", t, func() {
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

		action, err = UpdateAction(&UpdateActionRequest{
			ID:          action.ID.Hex(),
			Description: "This is a test action",
		})
		So(err, ShouldBeNil)
		So(action, ShouldNotBeNil)
		So(action.Name, ShouldEqual, "Test Action")
		So(action.Description, ShouldEqual, "This is a test action")
	})
}
