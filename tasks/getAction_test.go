package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAction(t *testing.T) {
	Convey("GetAction", t, func() {
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

		a, err := GetAction(&GetActionRequest{
			ActionID: action.ID.Hex(),
		})
		So(err, ShouldBeNil)
		So(a, ShouldNotBeNil)
		So(a, ShouldResemble, action)
	})
}
