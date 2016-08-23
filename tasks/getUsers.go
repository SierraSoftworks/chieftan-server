package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type GetUsersRequest struct {
}

func GetUsers(req *GetUsersRequest) ([]models.User, error) {
	var users []models.User
	if err := models.DB().Users().Find(&bson.M{}).Select(&bson.M{
		"tokens": 0,
	}).All(&users); err != nil {
		return nil, errors.ServerError()
	}

	return users, nil
}
