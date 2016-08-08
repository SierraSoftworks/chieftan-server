package tasks

import (
	"../models"
	"gopkg.in/mgo.v2"
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
			return nil, NewError(409, "Conflict", "There is already another user with this email address.")
		}

		return nil, NewError(500, "Server Error", "We encountered an issue creating this user.")
	}

	return &user, nil
}
