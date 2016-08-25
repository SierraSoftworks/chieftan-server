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
	Convey("Status API", t, func() {
		setUpTest()
		ts := httptest.NewServer(Router())
		defer ts.Close()

		Convey("/v1/status", func() {
			url := fmt.Sprintf("%s/v1/status", ts.URL)

			Convey("GET", func() {
				req, err := http.NewRequest("GET", url, nil)
				So(err, ShouldBeNil)

				res, err := http.DefaultClient.Do(req)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
				So(res.StatusCode, ShouldEqual, 200)

				var status models.Status
				dec := json.NewDecoder(res.Body)
				So(dec.Decode(&status), ShouldBeNil)

				So(status.StartedAt.Unix(), ShouldEqual, startedAt.Unix())
			})
		})
	})
}
