package api

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
)

func init() {
	Router().Path("/v1/projects").Methods("GET").Handler(girder.NewHandler(getProjects).RequireAuthentication(getUser).LogRequests())
	Router().Path("/v1/projects").Methods("POST").Handler(girder.NewHandler(createProject).RequireAuthentication(getUser).LogRequests())
}

func getProjects(c *girder.Context) (interface{}, error) {
	req := tasks.GetProjectsRequest{}
	projects, err := tasks.GetProjects(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return projects, nil
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
