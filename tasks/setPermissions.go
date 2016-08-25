package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type SetPermissionsRequest struct {
	UserID      string
	Permissions []string
}

func SetPermissions(req *SetPermissionsRequest) (*models.AuditLogContext, error) {
	user, err := GetUser(&GetUserRequest{
		ID: req.UserID,
	})

	if err != nil {
		return nil, formatError(err)
	}

	err = models.DB().Users().UpdateId(req.UserID, &bson.M{
		"$set": &bson.M{
			"permissions": req.Permissions,
		},
	})

	if err != nil {
		return nil, formatError(err)
	}

	return &models.AuditLogContext{
		User:    user.Summary(),
		Request: req,
	}, nil
}
