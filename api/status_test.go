package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStatus(t *testing.T) {
	Convey("/v1/status", t, func() {
		setUpTest()
		ts := httptest.NewServer(Router())
		defer ts.Close()

		Convey("GET", func() {
			req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/status", ts.URL), nil)
			So(err, ShouldBeNil)

			res, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.StatusCode, ShouldEqual, 200)

			var status models.Status
			dec := json.NewDecoder(res.Body)
			So(dec.Decode(&status), ShouldBeNil)

			So(status, ShouldResemble, models.Status{
				StartedAt: startedAt,
			})
		})
	})
}
