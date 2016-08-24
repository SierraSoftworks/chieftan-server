package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type RemoveTokenRequest struct {
	Token string
}

func RemoveToken(req *RemoveTokenRequest) (*models.AuditLogContext, error) {
	if !models.IsWellFormattedAccessToken(req.Token) {
		return nil, errors.BadRequest()
	}

	user, err := GetUser(&GetUserRequest{
		Token: req.Token,
	})
	if err != nil {
		return nil, formatError(err)
	}

	err = models.DB().Users().Update(&bson.M{
		"tokens": req.Token,
	}, &bson.M{
		"$pull": &bson.M{
			"tokens": req.Token,
		},
	})

	if err != nil {
		return nil, formatError(err)
	}

	return &models.AuditLogContext{
		User: user.Summary(),
	}, nil
}
