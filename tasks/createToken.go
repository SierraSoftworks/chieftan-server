package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CreateTokenRequest struct {
	UserID string
}

func CreateToken(req *CreateTokenRequest) (string, error) {
	if !models.IsValidUserID(req.UserID) {
		return "", errors.BadRequest()
	}

	token, err := models.GenerateAccessToken()
	if err != nil {
		return "", err
	}

	changes, err := models.DB().Users().UpdateAll(
		&bson.M{"_id": req.UserID},
		&bson.M{
			"$addToSet": &bson.M{
				"tokens": token,
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

	return token, nil
}
