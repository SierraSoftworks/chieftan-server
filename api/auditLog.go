package api

import (
	"github.com/SierraSoftworks/chieftan-server/tasks"
	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
)

func init() {
	Router().Path("/v1/audit").Methods("GET").Handler(girder.NewHandler(getAuditLog).RequireAuthentication(getUser).RequirePermission("admin").LogRequests())
	Router().Path("/v1/audit/{entry}").Methods("GET").Handler(girder.NewHandler(getAuditLogEntryByID).RequireAuthentication(getUser).RequirePermission("admin").LogRequests())
}

func getAuditLogEntryByID(c *girder.Context) (interface{}, error) {
	req := tasks.GetAuditLogEntryRequest{
		ID: c.Vars["entry"],
	}

	entry, err := tasks.GetAuditLogEntry(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return entry, nil
}

func getAuditLog(c *girder.Context) (interface{}, error) {
	req := tasks.GetAuditLogRequest{}

	entries, err := tasks.GetAuditLog(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return entries, nil
}