package api

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
)

func init() {
	Router().
		Path("/v1/project/{project}/actions").
		Methods("GET").
		Handler(girder.NewHandler(getActions).
			RequireAuthentication(getUser).
			RequirePermission("project/:project").
			LogRequests()).
		Name("GET /project/{project}/actions")
	Router().
		Path("/v1/project/{project}/actions").
		Methods("POST").
		Handler(girder.NewHandler(createAction).
			RequireAuthentication(getUser).
			RequirePermission("project/:project/admin").
			LogRequests()).
		Name("POST /project/{project}/actions")

	Router().
		Path("/v1/action/{action}").
		Methods("GET").
		Handler(girder.NewHandler(getAction).
			RequireAuthentication(getUser).
			LogRequests()).
		Name("GET /action/{action}")
	Router().
		Path("/v1/action/{action}").
		Methods("PUT").
		Handler(girder.NewHandler(updateAction).
			RequireAuthentication(getUser).
			LogRequests()).
		Name("PUT /action/{action}")
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

func getAction(c *girder.Context) (interface{}, error) {
	req := tasks.GetActionRequest{
		ActionID: c.Vars["action"],
	}

	action, err := tasks.GetAction(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	if !c.Permissions.WithContext(map[string]string{
		"project": action.Project.ID.Hex(),
	}).Can("project/:project") {
		return nil, errors.NotAllowed()
	}

	return action, nil
}

func updateAction(c *girder.Context) (interface{}, error) {
	action, err := tasks.GetAction(&tasks.GetActionRequest{
		ActionID: c.Vars["action"],
	})
	if err != nil {
		return nil, errors.From(err)
	}

	if !c.Permissions.WithContext(map[string]string{
		"project": action.Project.ID.Hex(),
	}).Can("project/:project/admin") {
		return nil, errors.NotAllowed()
	}

	req := tasks.UpdateActionRequest{}

	if err := c.ReadBody(&req); err != nil {
		return nil, errors.From(err)
	}

	req.ID = c.Vars["action"]

	action, err = tasks.UpdateAction(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return action, nil
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
