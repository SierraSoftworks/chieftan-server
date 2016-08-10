package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/src/models"
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
		return nil, NewError(500, "Server Error", "We encountered an error retrieving the users list from the database.")
	}

	if user.ID != req.UserID {
		return nil, NewError(404, "Not Found", "The user you specified could not be found in the database.")
	}

	return user.Tokens, nil
}
