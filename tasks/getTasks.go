package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type GetTasksRequest struct {
	ActionID  string `json:"action,omitempty"`
	ProjectID string `json:"project,omitempty"`
	Limit     int    `json:"limit,omitempty"`
	Skip      int    `json:"skip,omitempty"`
}

func GetTasks(req *GetTasksRequest) ([]models.Task, error) {
	query := bson.M{}

	if req.ActionID != "" {
		if !bson.IsObjectIdHex(req.ActionID) {
			return nil, errors.BadRequest()
		}

		query["action.id"] = bson.ObjectIdHex(req.ActionID)
	}

	if req.ProjectID != "" {
		if !bson.IsObjectIdHex(req.ProjectID) {
			return nil, errors.BadRequest()
		}

		query["project.id"] = bson.ObjectIdHex(req.ProjectID)
	}

	q := models.DB().Tasks().Find(&query).Sort("-created")

	if req.Skip != 0 {
		q = q.Skip(req.Skip)
	}

	if req.Limit != 0 {
		q = q.Limit(req.Limit)
	}

	var tasks []models.Task
	err := q.All(&tasks)
	if err != nil {
		return nil, formatError(err)
	}

	return tasks, nil
}
