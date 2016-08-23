package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type GetTokensRequest struct {
	UserID string
}

func GetTokens(req *GetTokensRequest) ([]string, error) {
	var user models.User
	if err := models.DB().Users().FindId(req.UserID).Select(&bson.M{
		"tokens": 1,
	}).One(&user); err != nil {
		return nil, errors.ServerError()
	}

	if user.ID != req.UserID {
		return nil, errors.NotFound()
	}

	return user.Tokens, nil
}
