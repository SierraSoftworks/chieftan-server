package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type GetProjectActionsRequest struct {
	ProjectID string `json:"project"`
}

func GetProjectActions(req *GetProjectActionsRequest) ([]models.Action, error) {
	if !bson.IsObjectIdHex(req.ProjectID) {
		return nil, errors.BadRequest()
	}

	var actions []models.Action
	err := models.DB().Actions().Find(&bson.M{
		"project.id": bson.ObjectIdHex(req.ProjectID),
	}).All(&actions)

	if err != nil {
		return nil, formatError(err)
	}

	return actions, nil
}
