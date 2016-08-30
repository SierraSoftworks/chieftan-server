package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type SetPermissionsRequest struct {
	UserID      string   `json:"user"`
	Permissions []string `json:"permissions"`
}

func SetPermissions(req *SetPermissionsRequest) (*models.User, *models.AuditLogContext, error) {
	user, err := GetUser(&GetUserRequest{
		UserID: req.UserID,
	})

	if err != nil {
		return nil, nil, formatError(err)
	}

	err = models.DB().Users().UpdateId(req.UserID, &bson.M{
		"$set": &bson.M{
			"permissions": req.Permissions,
		},
	})

	if err != nil {
		return nil, nil, formatError(err)
	}

	user.Permissions = req.Permissions

	return user, &models.AuditLogContext{
		User:    user.Summary(),
		Request: req,
	}, nil
}
