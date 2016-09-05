package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func testSetup() {
	So(Connect("mongodb://localhost/chieftan_test"), ShouldBeNil)

	_, err := DB().Users().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)

	_, err = DB().Projects().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)

	_, err = DB().Actions().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)

	_, err = DB().Tasks().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)

	_, err = DB().AuditLogs().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)
}
