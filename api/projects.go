package api

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
)

func init() {
	Router().
		Path("/v1/projects").
		Methods("GET").
		Handler(girder.NewHandler(getProjects).
			RequireAuthentication(getUser).
			LogRequests()).
		Name("GET /projects")
	Router().
		Path("/v1/projects").
		Methods("POST").
		Handler(girder.NewHandler(createProject).
			RequireAuthentication(getUser).
			LogRequests()).
		Name("POST /projects")

	Router().
		Path("/v1/project/{project}").
		Methods("GET").
		Handler(girder.NewHandler(getProject).
			RequireAuthentication(getUser).
			LogRequests()).
		Name("GET /project/{project}")
}

func getProjects(c *girder.Context) (interface{}, error) {
	req := tasks.GetProjectsRequest{}
	projects, err := tasks.GetProjects(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return projects, nil
}

func getProject(c *girder.Context) (interface{}, error) {
	req := tasks.GetProjectRequest{
		ProjectID: c.Vars["project"],
	}
	project, err := tasks.GetProject(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return project, nil
}

func createProject(c *girder.Context) (interface{}, error) {
	req := tasks.CreateProjectRequest{}
	if err := c.ReadBody(&req); err != nil {
		return nil, errors.From(err)
	}

	project, audit, err := tasks.CreateProject(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Type:    "project.create",
		User:    c.User.(*models.User).Summary(),
		Token:   c.GetAuthToken().Value,
		Context: audit,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return project, nil
}
