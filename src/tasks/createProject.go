package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/src/models"
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
		return nil, NewError(500, "Server Error", "We encountered an issue creating this project.")
	}

	return &project, nil
}
