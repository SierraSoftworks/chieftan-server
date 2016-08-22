package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
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
		return NewError(500, "Server Error", "We encountered an error removing the token from the database.")
	}

	return nil
}
