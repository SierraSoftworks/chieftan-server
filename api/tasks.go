package api

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
)

func init() {
	Router().Path("/v1/tasks").Methods("GET").Handler(girder.NewHandler(getTasks).RequireAuthentication(getUser).RequirePermission("admin").LogRequests())

	Router().Path("/v1/task/{task}").Methods("GET").Handler(girder.NewHandler(getTask).RequireAuthentication(getUser).LogRequests())
	Router().Path("/v1/task/{task}/run").Methods("POST").Handler(girder.NewHandler(runTask).RequireAuthentication(getUser).LogRequests())

	Router().Path("/v1/project/{project}/tasks").Methods("GET").Handler(girder.NewHandler(getProjectTasks).RequireAuthentication(getUser).RequirePermission("project/:project").LogRequests())
	Router().Path("/v1/project/{project}/tasks/recent").Methods("GET").Handler(girder.NewHandler(getRecentProjectTasks).RequireAuthentication(getUser).RequirePermission("project/:project").LogRequests())

	Router().Path("/v1/action/{action}/tasks").Methods("GET").Handler(girder.NewHandler(getActionTasks).RequireAuthentication(getUser).LogRequests())
	Router().Path("/v1/action/{action}/tasks").Methods("POST").Handler(girder.NewHandler(createTask).RequireAuthentication(getUser).LogRequests())
	Router().Path("/v1/action/{action}/tasks/recent").Methods("GET").Handler(girder.NewHandler(getRecentActionTasks).RequireAuthentication(getUser).LogRequests())
	Router().Path("/v1/action/{action}/task/latest").Methods("GET").Handler(girder.NewHandler(getLatestActionTask).RequireAuthentication(getUser).LogRequests())
	Router().Path("/v1/action/{action}/task/latest/run").Methods("POST").Handler(girder.NewHandler(runLatestActionTask).RequireAuthentication(getUser).LogRequests())
}

func getTasks(c *girder.Context) (interface{}, error) {
	tasks, err := tasks.GetTasks(&tasks.GetTasksRequest{})
	if err != nil {
		return nil, errors.From(err)
	}

	return tasks, nil
}

func getTask(c *girder.Context) (interface{}, error) {
	task, err := tasks.GetTask(&tasks.GetTaskRequest{
		TaskID: c.Vars["task"],
	})
	if err != nil {
		return nil, errors.From(err)
	}

	if !c.Permissions.WithContext(map[string]string{
		"project": task.Project.ID.Hex(),
	}).Can("project/:project") {
		return nil, errors.NotAllowed()
	}

	return task, nil
}

func runTask(c *girder.Context) (interface{}, error) {
	task, err := tasks.GetTask(&tasks.GetTaskRequest{
		TaskID: c.Vars["task"],
	})
	if err != nil {
		return nil, errors.From(err)
	}

	if !c.Permissions.WithContext(map[string]string{
		"project": task.Project.ID.Hex(),
	}).Can("project/:project") {
		return nil, errors.NotAllowed()
	}

	req := tasks.RunTaskRequest{}
	if err := c.ReadBody(&req); err != nil {
		return nil, errors.From(err)
	}

	req.TaskID = task.ID.Hex()

	exec, audit, err := tasks.RunTask(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Token:   c.GetAuthToken().Value,
		User:    c.User.(*models.User).Summary(),
		Type:    "task.run",
		Context: audit,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return exec.Task, nil
}

func getProjectTasks(c *girder.Context) (interface{}, error) {
	tasks, err := tasks.GetTasks(&tasks.GetTasksRequest{
		ProjectID: c.Vars["project"],
	})
	if err != nil {
		e := errors.From(err)
		if e.Code == 404 {
			return nil, errors.NotAllowed()
		}

		return nil, err
	}

	return tasks, nil
}

func getRecentProjectTasks(c *girder.Context) (interface{}, error) {
	tasks, err := tasks.GetTasks(&tasks.GetTasksRequest{
		ProjectID: c.Vars["project"],
		Limit:     50,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return tasks, nil
}

func getActionTasks(c *girder.Context) (interface{}, error) {
	action, err := tasks.GetAction(&tasks.GetActionRequest{
		ActionID: c.Vars["action"],
	})
	if err != nil {
		return nil, errors.From(err)
	}

	if !c.Permissions.WithContext(map[string]string{
		"project": action.Project.ID.Hex(),
	}).Can("project/:project") {
		return nil, errors.NotAllowed()
	}

	tasks, err := tasks.GetTasks(&tasks.GetTasksRequest{
		ActionID: action.ID.Hex(),
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return tasks, nil
}

func getRecentActionTasks(c *girder.Context) (interface{}, error) {
	action, err := tasks.GetAction(&tasks.GetActionRequest{
		ActionID: c.Vars["action"],
	})
	if err != nil {
		return nil, errors.From(err)
	}

	if !c.Permissions.WithContext(map[string]string{
		"project": action.Project.ID.Hex(),
	}).Can("project/:project") {
		return nil, errors.NotAllowed()
	}

	tasks, err := tasks.GetTasks(&tasks.GetTasksRequest{
		ActionID: action.ID.Hex(),
		Limit:    50,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return tasks, nil
}

func getLatestActionTask(c *girder.Context) (interface{}, error) {
	action, err := tasks.GetAction(&tasks.GetActionRequest{
		ActionID: c.Vars["action"],
	})
	if err != nil {
		return nil, errors.From(err)
	}

	if !c.Permissions.WithContext(map[string]string{
		"project": action.Project.ID.Hex(),
	}).Can("project/:project") {
		return nil, errors.NotAllowed()
	}

	tasks, err := tasks.GetTasks(&tasks.GetTasksRequest{
		ActionID: action.ID.Hex(),
		Limit:    1,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	if len(tasks) == 0 {
		return nil, errors.NotFound()
	}

	return tasks[0], nil
}

func runLatestActionTask(c *girder.Context) (interface{}, error) {
	action, err := tasks.GetAction(&tasks.GetActionRequest{
		ActionID: c.Vars["action"],
	})
	if err != nil {
		return nil, errors.From(err)
	}

	if !c.Permissions.WithContext(map[string]string{
		"project": action.Project.ID.Hex(),
	}).Can("project/:project") {
		return nil, errors.NotAllowed()
	}

	ts, err := tasks.GetTasks(&tasks.GetTasksRequest{
		ActionID: action.ID.Hex(),
		Limit:    1,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	if len(ts) == 0 {
		return nil, errors.NotFound()
	}

	task := ts[0]

	req := tasks.RunTaskRequest{}
	if err := c.ReadBody(&req); err != nil {
		return nil, errors.From(err)
	}

	req.TaskID = task.ID.Hex()

	exec, audit, err := tasks.RunTask(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Token:   c.GetAuthToken().Value,
		User:    c.User.(*models.User).Summary(),
		Type:    "task.run",
		Context: audit,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return exec.Task, nil
}

func createTask(c *girder.Context) (interface{}, error) {
	action, err := tasks.GetAction(&tasks.GetActionRequest{
		ActionID: c.Vars["action"],
	})
	if err != nil {
		return nil, errors.From(err)
	}

	req := tasks.CreateTaskRequest{}
	if err := c.ReadBody(&req); err != nil {
		return nil, errors.From(err)
	}

	req.Project = action.Project
	req.Action = action.Summary()

	task, audit, err := tasks.CreateTask(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Token:   c.GetAuthToken().Value,
		User:    c.User.(*models.User).Summary(),
		Type:    "task.create",
		Context: audit,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return task, nil
}
