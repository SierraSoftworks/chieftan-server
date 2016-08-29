package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
)

type RemoveUserRequest struct {
	UserID string `json:"user"`
}

func RemoveUser(req *RemoveUserRequest) (*models.User, *models.AuditLogContext, error) {
	var user models.User
	err := models.DB().Users().FindId(req.UserID).One(&user)
	if err != nil {
		return nil, nil, formatError(err)
	}

	err = models.DB().Users().RemoveId(req.UserID)
	if err != nil {
		return nil, nil, formatError(err)
	}

	return &user, &models.AuditLogContext{
		User:    user.Summary(),
		Request: req,
	}, nil
}
