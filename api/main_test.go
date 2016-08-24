package api

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

func setUpTest(t *testing.T) {
	models.Connect("mongodb://localhost/chieftan_test")

	_, err := models.DB().Users().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)

	_, err = models.DB().AuditLogs().RemoveAll(&bson.M{})
	So(err, ShouldBeNil)
}
