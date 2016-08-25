package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type CreateProjectRequest struct {
	Name        string
	Description string
	URL         string
}

func CreateProject(req *CreateProjectRequest) (*models.Project, *models.AuditLogContext, error) {
	project := models.Project{
		ID:          bson.NewObjectId(),
		Name:        req.Name,
		Description: req.Description,
		URL:         req.URL,
	}

	if err := models.DB().Projects().Insert(&project); err != nil {
		return nil, nil, formatError(err)
	}

	return &project, &models.AuditLogContext{
		Project: project.Summary(),
		Request: req,
	}, nil
}
