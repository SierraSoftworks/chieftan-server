package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateAction(t *testing.T) {
	Convey("CreateAction", t, func() {
		testSetup()

		project, _, err := CreateProject(&CreateProjectRequest{
			Name:        "Test Project",
			Description: "A test project",
			URL:         "http://test.com",
		})
		So(err, ShouldBeNil)
		So(project, ShouldNotBeNil)

		action, audit, err := CreateAction(&CreateActionRequest{
			Name:        "Test Action",
			Description: "A test action",
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
				Method: "GET",
				URL:    "https://github.com/SierraSoftworks/chieftan-server",
			},
			Project: project.Summary(),
		})
		So(err, ShouldBeNil)
		So(audit, ShouldNotBeNil)
		So(action, ShouldNotBeNil)

		So(audit.Project, ShouldResemble, project.Summary())
		So(audit.Action, ShouldResemble, action.Summary())

		So(action.Project, ShouldResemble, project.Summary())

		Convey("Updates database", func() {
			action, err := GetAction(&GetActionRequest{
				ActionID: action.ID.Hex(),
			})
			So(err, ShouldBeNil)
			So(action, ShouldNotBeNil)
		})
	})
}
