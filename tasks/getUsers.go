package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type GetUsersRequest struct {
}

func GetUsers(req *GetUsersRequest) ([]models.User, error) {
	users := []models.User{}
	err := models.DB().Users().Find(&bson.M{}).Select(&bson.M{
		"tokens": 0,
	}).All(&users)

	if err != nil {
		return nil, formatError(err)
	}

	return users, nil
}
