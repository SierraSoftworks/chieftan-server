package tasks

import (
	"../models"
	"gopkg.in/mgo.v2/bson"
)

type RemoveTokenRequest struct {
	Token string
}

func RemoveToken(req *RemoveTokenRequest) error {
	if !models.IsWellFormattedAccessToken(req.Token) {
		return NewError(400, "Bad Request", "The token you provided was not formatted correctly.")
	}

	changes, err := models.DB().Users().UpdateAll(&bson.M{"tokens": req.Token}, &bson.M{
		"$pull": &bson.M{
			"tokens": req.Token,
		},
	})
	if err != nil {
		return NewError(500, "Server Error", "We encountered an error removing the token from the database.")
	}

	if changes.Updated == 0 {
		return NewError(404, "Not Found", "The token you specified could not be found in the database.")
	}

	return nil
}
