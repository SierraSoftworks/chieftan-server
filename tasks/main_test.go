package tasks

import (
	"github.com/SierraSoftworks/chieftan-server/models"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func testSetup() {
	So(models.Connect("mongodb://localhost/chieftan_test"), ShouldBeNil)

	_, err := models.DB().Users().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)

	_, err = models.DB().Projects().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)

	_, err = models.DB().Actions().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)

	_, err = models.DB().Tasks().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)

	_, err = models.DB().AuditLogs().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)
}
