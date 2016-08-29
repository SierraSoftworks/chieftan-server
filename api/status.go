package api

import (
	"time"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder"
)

var startedAt = time.Now()

func init() {
	Router().
		Path("/v1/status").
		Methods("GET").
		Handler(girder.NewHandler(getStatus)).
		Name("GET /status")
}

func getStatus(c *girder.Context) (interface{}, error) {
	return &models.Status{
		StartedAt: startedAt,
	}, nil
}
