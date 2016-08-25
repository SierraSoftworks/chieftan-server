package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type GetProjectRequest struct {
	ProjectID string `json:"project"`
}

func GetProject(req *GetProjectRequest) (*models.Project, error) {
	if !bson.IsObjectIdHex(req.ProjectID) {
		return nil, errors.BadRequest()
	}

	var project models.Project
	err := models.DB().Projects().FindId(bson.ObjectIdHex(req.ProjectID)).One(&project)
	if err != nil {
		return nil, formatError(err)
	}

	return &project, nil
}
