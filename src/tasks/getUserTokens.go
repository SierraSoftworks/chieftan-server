package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/src/models"
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
		return nil, NewError(500, "Server Error", "We encountered an error retrieving the users list from the database.")
	}

	return user.Tokens, nil
}
