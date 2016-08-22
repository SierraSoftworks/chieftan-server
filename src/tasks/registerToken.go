package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/src/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegisterTokenRequest struct {
	UserID string
	Token  string
}

func RegisterToken(req *RegisterTokenRequest) (string, error) {
	if !models.IsValidUserID(req.UserID) {
		return "", NewError(400, "Bad Request", "You failed to provide a well formed user ID.")
	}

	if !models.IsWellFormattedAccessToken(req.Token) {
		return "", NewError(400, "Bad Request", "You failed to provide a well formed access token.")
	}

	changes, err := models.DB().Users().UpdateAll(
		&bson.M{"_id": req.UserID},
		&bson.M{
			"$addToSet": &bson.M{
				"tokens": req.Token,
			},
		})
	if err != nil {
		if mgo.IsDup(err) {
			return "", NewError(409, "Conflict", "The generated access token has already been used, please try again.")
		}

		return "", NewError(500, "Server Error", "We encountered an error updating the database with the generated access token.")
	}

	if changes.Updated == 0 {
		return "", NewError(404, "Not Found", "The user you specified could not be found in the database.")
	}

	return req.Token, nil
}
