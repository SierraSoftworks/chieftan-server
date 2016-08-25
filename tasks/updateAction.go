package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type UpdateActionRequest struct {
	ID             string                       `json:"action"`
	Name           string                       `json:"name,omitempty"`
	Description    string                       `json:"description,omitempty"`
	Variables      map[string]string            `json:"vars,omitempty"`
	HTTP           *models.Request              `json:"http,omitempty"`
	Configurations []models.ActionConfiguration `json:"configurations,omitempty"`
}

func UpdateAction(req *UpdateActionRequest) (*models.Action, error) {
	if !bson.IsObjectIdHex(req.ID) {
		return nil, errors.BadRequest()
	}

	var action models.Action
	err := models.DB().Actions().FindId(bson.ObjectIdHex(req.ID)).One(&action)
	if err != nil {
		return nil, formatError(err)
	}

	if req.Name != "" {
		action.Name = req.Name
	}

	if req.Description != "" {
		action.Description = req.Description
	}

	if req.Variables != nil {
		action.Variables = req.Variables
	}

	if req.HTTP != nil {
		action.HTTP = req.HTTP
	}

	if req.Configurations != nil {
		action.Configurations = req.Configurations
	}

	err = models.DB().Actions().UpdateId(bson.ObjectIdHex(req.ID), action)
	if err != nil {
		return nil, formatError(err)
	}

	return &action, nil
}
