package tasks

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateProject(t *testing.T) {
	Convey("CreateProject", t, func() {
		testSetup()

		req := &CreateProjectRequest{
			Name:        "Test Project",
			Description: "A test project",
			URL:         "http://test.com",
		}

		project, audit, err := CreateProject(req)
		So(err, ShouldBeNil)
		So(audit, ShouldNotBeNil)
		So(audit.Project, ShouldNotBeNil)
		So(audit.Project.ID, ShouldEqual, project.ID)
		So(audit.Project.Name, ShouldEqual, "Test Project")

		So(project, ShouldNotBeNil)
		So(project.ID, ShouldNotEqual, "")
		So(project.Name, ShouldEqual, "Test Project")
		So(project.Description, ShouldEqual, "A test project")
		So(project.URL, ShouldEqual, "http://test.com")
	})
}
