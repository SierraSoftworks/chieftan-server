package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type RemoveTokenRequest struct {
	Token string
}

func RemoveToken(req *RemoveTokenRequest) error {
	if !models.IsWellFormattedAccessToken(req.Token) {
		return errors.BadRequest()
	}

	changes, err := models.DB().Users().UpdateAll(&bson.M{"tokens": req.Token}, &bson.M{
		"$pull": &bson.M{
			"tokens": req.Token,
		},
	})
	if err != nil {
		return errors.ServerError()
	}

	if changes.Updated == 0 {
		return errors.NotFound()
	}

	return nil
}
