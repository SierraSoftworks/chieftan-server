package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type RemovePermissionsRequest struct {
	UserID      string   `json:"user"`
	Permissions []string `json:"permissions"`
}

func RemovePermissions(req *RemovePermissionsRequest) (*models.User, *models.AuditLogContext, error) {
	user, err := GetUser(&GetUserRequest{
		ID: req.UserID,
	})

	if err != nil {
		return nil, nil, formatError(err)
	}

	err = models.DB().Users().UpdateId(req.UserID, &bson.M{
		"$pullAll": &bson.M{
			"permissions": req.Permissions,
		},
	})
	if err != nil {
		return nil, nil, formatError(err)
	}

	user, err = GetUser(&GetUserRequest{
		ID: req.UserID,
	})

	if err != nil {
		return nil, nil, formatError(err)
	}

	return user, &models.AuditLogContext{
		User:    user.Summary(),
		Request: req,
	}, nil
}
