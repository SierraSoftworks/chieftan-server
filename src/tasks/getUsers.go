package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/src/models"
	"gopkg.in/mgo.v2/bson"
)

type GetUsersRequest struct {
}

func GetUsers(req *GetUsersRequest) ([]models.User, error) {
	var users []models.User
	if err := models.DB().Users().Find(&bson.M{}).Select(&bson.M{
		"tokens": 0,
	}).All(&users); err != nil {
		return nil, NewError(500, "Server Error", "We encountered an error retrieving the users list from the database.")
	}

	return users, nil
}
