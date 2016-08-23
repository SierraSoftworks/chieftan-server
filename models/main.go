package models

import (
	"log"
	"net/url"
	"strings"

	"gopkg.in/mgo.v2"
)

var db Database

func DB() *Database {
	return &db
}

func Connect(mongodbURL string) {
	urlInfo, err := url.Parse(mongodbURL)
	if err != nil {
		log.Fatal(err)
	}

	session, err := mgo.Dial(urlInfo.Host)

	if err != nil {
		log.Fatal(err)
	}

	db = Database{
		db:      session.DB(strings.TrimLeft(urlInfo.Path, "/")),
		session: session,
	}
}
