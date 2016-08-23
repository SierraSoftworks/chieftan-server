package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type CreateProjectRequest struct {
	Name        string
	Description string
	URL         string
}

func CreateProject(req *CreateProjectRequest) (*models.Project, error) {
	project := models.Project{
		ID:          string(bson.NewObjectId()),
		Name:        req.Name,
		Description: req.Description,
		URL:         req.URL,
	}

	if err := models.DB().Projects().Insert(&project); err != nil {
		return nil, errors.ServerError()
	}

	return &project, nil
}
