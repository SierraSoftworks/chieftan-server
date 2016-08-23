package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	mgo "gopkg.in/mgo.v2"
)

type CreateUserRequest struct {
	Name        string
	Email       string
	Permissions []string
}

func CreateUser(req *CreateUserRequest) (*models.User, error) {
	user := models.User{
		Name:        req.Name,
		Email:       req.Email,
		Permissions: req.Permissions,
		Tokens:      []string{},
	}

	user.UpdateID()

	if err := models.DB().Users().Insert(&user); err != nil {
		if mgo.IsDup(err) {
			return nil, errors.Conflict()
		}

		return nil, errors.ServerError()
	}

	return &user, nil
}
