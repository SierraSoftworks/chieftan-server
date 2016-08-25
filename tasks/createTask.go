package tasks

import (
	"time"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type CreateTaskRequest struct {
	Metadata  *models.TaskMetadata   `json:"metadata"`
	Variables map[string]string      `json:"vars"`
	Action    *models.ActionSummary  `json:"-"`
	Project   *models.ProjectSummary `json:"-"`
}

func CreateTask(req *CreateTaskRequest) (*models.Task, *models.AuditLogContext, error) {
	if req.Project == nil {
		return nil, nil, errors.BadRequest()
	}

	if req.Action == nil {
		return nil, nil, errors.BadRequest()
	}

	if req.Metadata == nil {
		return nil, nil, errors.BadRequest()
	}

	task := models.Task{
		ID:        bson.NewObjectId(),
		Metadata:  *req.Metadata,
		Action:    req.Action,
		Project:   req.Project,
		Variables: req.Variables,
		Created:   time.Now(),
		State:     models.TaskStateNotExecuted,
	}

	if err := models.DB().Tasks().Insert(&task); err != nil {
		return nil, nil, formatError(err)
	}

	return &task, &models.AuditLogContext{
		Action:  req.Action,
		Project: req.Project,
		Task:    task.Summary(),
		Request: req,
	}, nil
}
