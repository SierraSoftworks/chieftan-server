package tasks

import (
	"github.com/SierraSoftworks/girder/errors"
	log "github.com/Sirupsen/logrus"
	raven "github.com/getsentry/raven-go"
	"gopkg.in/mgo.v2"
)

func formatError(err error) error {
	if err == nil {
		return nil
	}

	if mgo.IsDup(err) {
		return errors.Conflict()
	}

	if err.Error() == "not found" {
		return errors.NotFound()
	}

	log.Error(err)
	raven.CaptureError(err, nil)
	return errors.From(err)
}
