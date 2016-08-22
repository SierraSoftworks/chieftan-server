package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/src/models"
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
		return nil, NewError(500, "Server Error", "We encountered an error updating this user's permissions.")
	}

	return GetUser(&GetUserRequest{ID: req.UserID})
}
