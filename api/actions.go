package api

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
)

func init() {
	Router().Path("/v1/project/{project}/actions").Methods("GET").Handler(girder.NewHandler(getActions).RequireAuthentication(getUser).LogRequests())
	Router().Path("/v1/project/{project}/actions").Methods("POST").Handler(girder.NewHandler(createAction).RequireAuthentication(getUser).LogRequests())
}

func getActions(c *girder.Context) (interface{}, error) {
	req := tasks.GetProjectActionsRequest{
		ProjectID: c.Vars["project"],
	}

	actions, err := tasks.GetProjectActions(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return actions, nil
}

func createAction(c *girder.Context) (interface{}, error) {
	project, err := tasks.GetProject(&tasks.GetProjectRequest{
		ProjectID: c.Vars["project"],
	})

	if err != nil {
		return nil, errors.From(err)
	}

	req := tasks.CreateActionRequest{}
	if err := c.ReadBody(&req); err != nil {
		return nil, errors.From(err)
	}

	req.Project = project.Summary()

	action, audit, err := tasks.CreateAction(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Type:    "action.create",
		User:    c.User.(*models.User).Summary(),
		Token:   c.GetAuthToken().Value,
		Context: audit,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return action, nil
}
