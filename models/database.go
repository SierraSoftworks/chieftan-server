package models

import (
	"gopkg.in/mgo.v2"
)

type Database struct {
	db *mgo.Database
}

func (d *Database) Actions() *mgo.Collection {
	return d.db.C("actions")
}

func (d *Database) Projects() *mgo.Collection {
	return d.db.C("projects")
}

func (d *Database) Tasks() *mgo.Collection {
	return d.db.C("tasks")
}

func (d *Database) AuditLogs() *mgo.Collection {
	return d.db.C("auditLogs")
}

func (d *Database) Users() *mgo.Collection {
	return d.db.C("users")
}
