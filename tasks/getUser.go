package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type GetUserRequest struct {
	ID    string `json:"user"`
	Token string `json:"token"`
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

	err := models.DB().Users().Find(&query).Select(&bson.M{
		"tokens": 0,
	}).One(&user)

	if err != nil {
		return nil, formatError(err)
	}

	return &user, nil
}
