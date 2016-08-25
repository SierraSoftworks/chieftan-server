package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type CreateTokenRequest struct {
	UserID string `json:"user"`
}

func CreateToken(req *CreateTokenRequest) (string, *models.AuditLogContext, error) {
	if !models.IsValidUserID(req.UserID) {
		return "", nil, errors.BadRequest()
	}

	user, err := GetUser(&GetUserRequest{ID: req.UserID})
	if err != nil {
		return "", nil, err
	}

	token, err := models.GenerateAccessToken()
	if err != nil {
		return "", nil, err
	}

	err = models.DB().Users().Update(
		&bson.M{"_id": req.UserID},
		&bson.M{
			"$addToSet": &bson.M{
				"tokens": token,
			},
		})

	if err != nil {
		return "", nil, formatError(err)
	}

	return token, &models.AuditLogContext{
		User:    user.Summary(),
		Request: req,
	}, nil
}
