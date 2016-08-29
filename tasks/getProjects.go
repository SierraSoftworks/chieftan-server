package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type GetProjectsRequest struct {
}

func GetProjects(req *GetProjectsRequest) ([]models.Project, error) {
	projects := []models.Project{}
	err := models.DB().Projects().Find(&bson.M{}).All(&projects)
	if err != nil {
		return nil, formatError(err)
	}

	return projects, nil
}
