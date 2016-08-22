package api

import "gopkg.in/mgo.v2/bson"
import "github.com/SierraSoftworks/chieftan-server/models"
import "github.com/SierraSoftworks/chieftan-server/api/utils"

type userStore struct {
}

type authenticatedUser struct {
	id          string
	permissions []string
}

func (u *authenticatedUser) ID() string {
	return u.id
}

func (u *authenticatedUser) Permissions() []string {
	return u.permissions
}

func (s *userStore) GetUser(token *utils.AuthorizationToken) (utils.User, *utils.Error) {
	var user models.User
	err := models.DB().Users().Find(&bson.M{
		"tokens": token.Value,
	}).Select(&bson.M{
		"_id":         1,
		"permissions": 1,
	}).One(&user)

	if err != nil {
		return nil, utils.NewErrorFor(err)
	}

	if user.ID == "" {
		return nil, nil
	}

	return &authenticatedUser{
		user.ID,
		user.Permissions,
	}, nil
}
