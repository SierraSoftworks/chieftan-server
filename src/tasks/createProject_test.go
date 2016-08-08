package tasks

import . "gopkg.in/check.v1"

func (s *TasksSuite) TestCreateProject(c *C) {
	req := &CreateProjectRequest{
		Name:        "Test Project",
		Description: "A test project",
		URL:         "http://test.com",
	}

	project, err := CreateProject(req)
	c.Assert(err, IsNil)
	c.Assert(project, NotNil)

	c.Check(project.ID, Not(Equals), "")
	c.Check(project.Name, Equals, "Test Project")
	c.Check(project.Description, Equals, "A test project")
	c.Check(project.URL, Equals, "http://test.com")
}