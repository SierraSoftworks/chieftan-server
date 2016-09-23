package models

import (
	"fmt"
	"net/url"
	"strings"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

var db Database

func DB() *Database {
	return &db
}

func Connect(mongodbURL string) error {
	urlInfo, err := url.Parse(mongodbURL)
	if err != nil {
		return fmt.Errorf("failed to parse url: %s", err)
	}

	log.WithField("URL", urlInfo.Host).Debug("connecting");
	session, err := mgo.Dial(urlInfo.Host)
	if err != nil {
		return fmt.Errorf("failed to connect: %s", err)
	}

	db = Database{
		db:      session.DB(strings.TrimLeft(urlInfo.Path, "/")),
		session: session,
	}

	return nil
}
