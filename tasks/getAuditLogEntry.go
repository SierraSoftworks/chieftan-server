package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
	"gopkg.in/mgo.v2/bson"
)

type GetAuditLogEntryRequest struct {
	ID string `json:"entry"`
}

func GetAuditLogEntry(req *GetAuditLogEntryRequest) (*models.AuditLog, error) {
	var auditLogEntry models.AuditLog

	if !bson.IsObjectIdHex(req.ID) {
		return nil, errors.BadRequest()
	}

	err := models.DB().AuditLogs().FindId(bson.ObjectIdHex(req.ID)).One(&auditLogEntry)

	if err != nil {
		return nil, formatError(err)
	}

	return &auditLogEntry, nil
}
