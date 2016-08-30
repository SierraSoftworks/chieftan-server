package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type RegisterTokenRequest struct {
	UserID string `json:"user"`
	Token  string `json:"token"`
}

func RegisterToken(req *RegisterTokenRequest) (string, *models.AuditLogContext, error) {
	if !models.IsValidUserID(req.UserID) {
		return "", nil, errors.BadRequest()
	}

	if !models.IsWellFormattedAccessToken(req.Token) {
		return "", nil, errors.BadRequest()
	}

	user, err := GetUser(&GetUserRequest{UserID: req.UserID})
	if err != nil {
		return "", nil, err
	}

	err = models.DB().Users().Update(
		&bson.M{"_id": req.UserID},
		&bson.M{
			"$addToSet": &bson.M{
				"tokens": req.Token,
			},
		})

	if err != nil {
		return "", nil, formatError(err)
	}

	return req.Token, &models.AuditLogContext{
		User:    user.Summary(),
		Request: req,
	}, nil
}
