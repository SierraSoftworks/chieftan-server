package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type RemoveAllTokensRequest struct {
	UserID string
}

func RemoveAllTokens(req *RemoveAllTokensRequest) (*models.AuditLogContext, error) {
	query := bson.M{}

	auditContext := models.AuditLogContext{}

	if req.UserID != "" {
		query["_id"] = req.UserID

		user, err := GetUser(&GetUserRequest{ID: req.UserID})
		if err != nil {
			return nil, err
		}

		auditContext.User = user.Summary()
	}

	_, err := models.DB().Users().UpdateAll(&query, &bson.M{
		"$set": &bson.M{
			"tokens": []string{},
		},
	})

	return &auditContext, formatError(err)
}
