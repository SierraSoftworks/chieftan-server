package tasks

import (
	"github.com/SierraSoftworks/girder/errors"
	log "github.com/Sirupsen/logrus"
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
	return errors.From(err)
}
