package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegisterTokenRequest struct {
	UserID string
	Token  string
}

func RegisterToken(req *RegisterTokenRequest) (string, error) {
	if !models.IsValidUserID(req.UserID) {
		return "", errors.BadRequest()
	}

	if !models.IsWellFormattedAccessToken(req.Token) {
		return "", errors.BadRequest()
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
			return "", errors.Conflict()
		}

		return "", errors.ServerError()
	}

	if changes.Updated == 0 {
		return "", errors.NotFound()
	}

	return req.Token, nil
}
