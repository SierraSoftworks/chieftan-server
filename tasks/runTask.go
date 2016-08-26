package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/executors"
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
)

type RunTaskRequest struct {
	TaskID        string            `json:"task"`
	Configuration string            `json:"configuration"`
	Variables     map[string]string `json:"vars"`
}

func RunTask(req *RunTaskRequest) (*executors.Execution, *models.AuditLogContext, error) {
	task, err := GetTask(&GetTaskRequest{
		TaskID: req.TaskID,
	})
	if err != nil {
		return nil, nil, formatError(err)
	}

	action, err := GetAction(&GetActionRequest{
		ActionID: task.Action.ID.Hex(),
	})
	if err != nil {
		return nil, nil, formatError(err)
	}

	var config *models.ActionConfiguration
	if req.Configuration != "" {
		for _, conf := range action.Configurations {
			if conf.Name == req.Configuration {
				config = &conf
				break
			}
		}

		if config == nil {
			return nil, nil, errors.NotFound()
		}
	}

	execution, err := executors.NewExecution(&executors.Options{
		Action:        action,
		Task:          task,
		Configuration: config,
		Variables:     req.Variables,
	})
	if err != nil {
		return nil, nil, formatError(err)
	}

	stateChanged := execution.Start()

	go func() {
		for task := range stateChanged {
			models.DB().Tasks().UpdateId(task.ID, task)
		}
	}()

	return execution, &models.AuditLogContext{
		Project: task.Project,
		Action:  task.Action,
		Task:    task.Summary(),
	}, nil
}
