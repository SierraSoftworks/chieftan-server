package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type GetTaskRequest struct {
	TaskID string `json:"task"`
}

func GetTask(req *GetTaskRequest) (*models.Task, error) {
	if !bson.IsObjectIdHex(req.TaskID) {
		return nil, errors.BadRequest()
	}

	var task models.Task
	err := models.DB().Tasks().FindId(bson.ObjectIdHex(req.TaskID)).One(&task)

	if err != nil {
		return nil, formatError(err)
	}

	return &task, nil
}
