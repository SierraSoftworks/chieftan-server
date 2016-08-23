package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type RemoveAllTokensRequest struct {
	UserID string
}

func RemoveAllTokens(req *RemoveAllTokensRequest) error {
	query := bson.M{}

	if req.UserID != "" {
		query["_id"] = req.UserID
	}

	_, err := models.DB().Users().UpdateAll(&query, &bson.M{
		"$set": &bson.M{
			"tokens": []string{},
		},
	})

	if err != nil {
		return errors.ServerError()
	}

	return nil
}
