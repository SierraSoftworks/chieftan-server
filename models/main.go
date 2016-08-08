package models

import (
	"log"
	"net/url"
	"os"
	"strings"

	"gopkg.in/mgo.v2"
)

var db Database

func DB() *Database {
	return &db
}

func init() {
	mongodbURL := os.ExpandEnv("$MONGODB_URL")
	if mongodbURL == "" {
		mongodbURL = "mongodb://localhost/chieftan"
	}

	urlInfo, err := url.Parse(mongodbURL)
	if err != nil {
		log.Fatal(err)
	}

	session, err := mgo.Dial(urlInfo.Host)

	if err != nil {
		log.Fatal(err)
	}

	db = Database{
		db: session.DB(strings.TrimLeft(urlInfo.Path, "/")),
	}
}
