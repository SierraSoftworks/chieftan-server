package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type GetActionRequest struct {
	ActionID string `json:"action"`
}

func GetAction(req *GetActionRequest) (*models.Action, error) {
	if !bson.IsObjectIdHex(req.ActionID) {
		return nil, errors.BadRequest()
	}

	var action models.Action
	err := models.DB().Actions().FindId(bson.ObjectIdHex(req.ActionID)).One(&action)
	if err != nil {
		return nil, formatError(err)
	}

	return &action, nil
}
