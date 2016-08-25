package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type CreateActionRequest struct {
	Name           string                       `json:"name"`
	Description    string                       `json:"description"`
	Variables      map[string]string            `json:"vars"`
	Configurations []models.ActionConfiguration `json:"configurations"`
	HTTP           *models.Request              `json:"http"`
	Project        *models.ProjectSummary       `json:"-"`
}

func CreateAction(req *CreateActionRequest) (*models.Action, *models.AuditLogContext, error) {
	if req.Project == nil {
		return nil, nil, errors.BadRequest()
	}

	action := models.Action{
		ID:             bson.NewObjectId(),
		Name:           req.Name,
		Description:    req.Description,
		Variables:      req.Variables,
		Configurations: req.Configurations,
		HTTP:           req.HTTP,
		Project:        req.Project,
	}

	if err := models.DB().Actions().Insert(&action); err != nil {
		return nil, nil, formatError(err)
	}

	return &action, &models.AuditLogContext{
		Action:  action.Summary(),
		Project: req.Project,
		Request: req,
	}, nil
}
