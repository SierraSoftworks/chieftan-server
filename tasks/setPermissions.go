package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type SetPermissionsRequest struct {
	UserID      string
	Permissions []string
}

func SetPermissions(req *SetPermissionsRequest) (*models.User, error) {
	if err := models.DB().Users().UpdateId(req.UserID, &bson.M{
		"$set": &bson.M{
			"permissions": req.Permissions,
		},
	}); err != nil {
		return nil, errors.ServerError()
	}

	return GetUser(&GetUserRequest{ID: req.UserID})
}
