package models

import (
	"fmt"
	"net/url"
	"strings"

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
