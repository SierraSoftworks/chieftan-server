package tasks

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/SierraSoftworks/chieftan-server/models"
)

type CreateAuditLogEntryRequest struct {
	Type    string                  `json:"type"`
	User    *models.UserSummary     `json:"user"`
	Token   string                  `json:"token"`
	Context *models.AuditLogContext `json:"context"`
}

func CreateAuditLogEntry(req *CreateAuditLogEntryRequest) (*models.AuditLog, error) {
	entry := models.AuditLog{
		ID:        bson.NewObjectId(),
		Type:      req.Type,
		User:      *req.User,
		Token:     req.Token,
		Timestamp: time.Now(),
		Context:   *req.Context,
	}

	if err := models.DB().AuditLogs().Insert(&entry); err != nil {
		return nil, formatError(err)
	}

	return &entry, nil
}
