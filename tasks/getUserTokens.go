package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type GetUserTokensRequest struct {
	ID string `json:"user"`
}

func GetUserTokens(req *GetUserTokensRequest) ([]string, *models.AuditLogContext, error) {
	var user models.User

	err := models.DB().Users().Find(&bson.M{}).Select(&bson.M{
		"_id":    1,
		"name":   1,
		"email":  1,
		"tokens": 1,
	}).One(&user)

	if err != nil {
		return nil, nil, formatError(err)
	}

	return user.Tokens, &models.AuditLogContext{
		User:    user.Summary(),
		Request: req,
	}, nil
}
