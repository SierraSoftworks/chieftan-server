package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type GetUserRequest struct {
	ID    string
	Token string
}

func GetUser(req *GetUserRequest) (*models.User, error) {
	var user models.User

	query := bson.M{}

	if req.ID != "" {
		query["_id"] = req.ID
	}

	if req.Token != "" {
		query["tokens"] = req.Token
	}

	if err := models.DB().Users().Find(&query).Select(&bson.M{
		"tokens": 0,
	}).One(&user); err != nil {
		return nil, NewError(500, "Server Error", "We encountered an error retrieving the users list from the database.")
	}

	if user.ID == "" {
		return nil, nil
	}

	return &user, nil
}
