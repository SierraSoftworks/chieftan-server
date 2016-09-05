package models

import (
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
		return err
	}

	session, err := mgo.Dial(urlInfo.Host)
	if err != nil {
		return err
	}

	db = Database{
		db:      session.DB(strings.TrimLeft(urlInfo.Path, "/")),
		session: session,
	}

	return nil
}
