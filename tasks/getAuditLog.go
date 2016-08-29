package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"gopkg.in/mgo.v2/bson"
)

type GetAuditLogRequest struct {
}

func GetAuditLog(req *GetAuditLogRequest) ([]models.AuditLog, error) {
	entries := []models.AuditLog{}

	if err := models.DB().AuditLogs().Find(&bson.M{}).All(&entries); err != nil {
		return nil, formatError(err)
	}

	return entries, nil
}
