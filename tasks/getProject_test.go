package tasks

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetProject(t *testing.T) {
	Convey("GetProject", t, func() {
		testSetup()

		newProject, _, err := CreateProject(&CreateProjectRequest{
			Name:        "Test Project",
			Description: "Test",
			URL:         "https://github.com/sierrasoftworks/chieftan-server",
		})
		So(err, ShouldBeNil)
		So(newProject, ShouldNotBeNil)

		project, err := GetProject(&GetProjectRequest{
			ProjectID: newProject.ID.Hex(),
		})
		So(err, ShouldBeNil)
		So(project, ShouldNotBeNil)
		So(project, ShouldResemble, newProject)
	})
}
