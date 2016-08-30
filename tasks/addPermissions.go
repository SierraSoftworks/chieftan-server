package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type AddPermissionsRequest struct {
	UserID      string   `json:"user"`
	Permissions []string `json:"permissions"`
}

func AddPermissions(req *AddPermissionsRequest) (*models.User, *models.AuditLogContext, error) {
	user, err := GetUser(&GetUserRequest{
		UserID: req.UserID,
	})

	if err != nil {
		return nil, nil, formatError(err)
	}

	err = models.DB().Users().UpdateId(req.UserID, &bson.M{
		"$addToSet": &bson.M{
			"permissions": &bson.M{
				"$each": req.Permissions,
			},
		},
	})
	if err != nil {
		return nil, nil, formatError(err)
	}

	user, err = GetUser(&GetUserRequest{
		UserID: req.UserID,
	})

	if err != nil {
		return nil, nil, formatError(err)
	}

	return user, &models.AuditLogContext{
		User:    user.Summary(),
		Request: req,
	}, nil
}
