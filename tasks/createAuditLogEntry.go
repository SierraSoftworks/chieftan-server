package tasks

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
)

type CreateAuditLogEntryRequest struct {
	Type    string
	User    models.UserSummary
	Token   string
	Context models.AuditLogContext
}

func CreateAuditLogEntry(req *CreateAuditLogEntryRequest) (*models.AuditLog, error) {
	entry := models.AuditLog{
		ID:        string(bson.NewObjectId()),
		Type:      req.Type,
		User:      req.User,
		Token:     req.Token,
		Timestamp: time.Now(),
		Context:   req.Context,
	}

	if err := models.DB().AuditLogs().Insert(&entry); err != nil {
		return nil, errors.ServerError()
	}

	return &entry, nil
}
