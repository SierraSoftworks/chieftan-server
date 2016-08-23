package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type GetUserTokensRequest struct {
	ID string
}

func GetUserTokens(req *GetUserTokensRequest) ([]string, error) {
	var user models.User
	if err := models.DB().Users().Find(&bson.M{}).Select(&bson.M{
		"tokens": 1,
	}).One(&user); err != nil {
		return nil, errors.ServerError()
	}

	return user.Tokens, nil
}
