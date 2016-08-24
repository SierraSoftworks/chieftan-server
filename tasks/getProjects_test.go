package tasks

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetProjects(t *testing.T) {
	Convey("GetProjects", t, func() {
		testSetup()

		newProject, _, err := CreateProject(&CreateProjectRequest{
			Name:        "Test Project",
			Description: "Test",
			URL:         "https://github.com/sierrasoftworks/chieftan-server",
		})
		So(err, ShouldBeNil)
		So(newProject, ShouldNotBeNil)

		projects, err := GetProjects(&GetProjectsRequest{})
		So(err, ShouldBeNil)
		So(projects, ShouldNotBeNil)
		So(projects, ShouldHaveLength, 1)
		So(projects, ShouldResemble, []models.Project{*newProject})
	})
}
