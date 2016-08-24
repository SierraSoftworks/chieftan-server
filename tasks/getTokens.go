package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type GetTokensRequest struct {
	UserID string
}

func GetTokens(req *GetTokensRequest) ([]string, *models.AuditLogContext, error) {
	var user models.User

	err := models.DB().Users().FindId(req.UserID).Select(&bson.M{
		"_id":    1,
		"name":   1,
		"email":  1,
		"tokens": 1,
	}).One(&user)

	if err != nil {
		return nil, nil, formatError(err)
	}

	return user.Tokens, &models.AuditLogContext{
		User: user.Summary(),
	}, nil
}
