package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	"github.com/SierraSoftworks/girder/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAuditLog(t *testing.T) {
	Convey("Audit Log API", t, func() {
		setUpTest()
		ts := httptest.NewServer(Router())
		defer ts.Close()

		entry, err := tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
			Type: "test",
			User: &models.UserSummary{
				ID:    "test",
				Name:  "Test User",
				Email: "test@test.com",
			},
			Token:   "0123456789abcdef0123456789abcdef",
			Context: &models.AuditLogContext{},
		})
		So(entry, ShouldNotBeNil)
		So(err, ShouldBeNil)

		Convey("/v1/audit", func() {

			Convey("GET", func() {
				req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/audit", ts.URL), nil)
				So(err, ShouldBeNil)

				Convey("When not signed in", func() {
					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 401)

					dec := json.NewDecoder(res.Body)
					var e errors.Error
					So(dec.Decode(&e), ShouldBeNil)
					So(e.Code, ShouldEqual, 401)
				})

				Convey("When signed in", func() {
					user, _, err := tasks.CreateUser(&tasks.CreateUserRequest{
						Name:  "Test User",
						Email: "test@test.com",
					})
					So(err, ShouldBeNil)
					So(user, ShouldNotBeNil)

					token, _, err := tasks.CreateToken(&tasks.CreateTokenRequest{
						UserID: user.ID,
					})
					So(err, ShouldBeNil)

					req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

					Convey("Without admin permissions", func() {
						res, err := http.DefaultClient.Do(req)
						defer res.Body.Close()
						So(err, ShouldBeNil)

						So(res.StatusCode, ShouldEqual, 403)
					})

					Convey("With admin permissions", func() {
						_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{"admin"},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						defer res.Body.Close()
						So(err, ShouldBeNil)

						So(res.StatusCode, ShouldEqual, 200)

						var entries []models.AuditLog
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&entries), ShouldBeNil)
						So(entries, ShouldHaveLength, 1)

						e := entries[0]
						So(e.ID, ShouldEqual, entry.ID)
						So(e.Timestamp.Unix(), ShouldEqual, entry.Timestamp.Unix())
						So(e.Token, ShouldEqual, entry.Token)
					})
				})
			})
		})

		Convey("/v1/audit/{entry}", func() {

			Convey("GET", func() {
				req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/audit/%s", ts.URL, entry.ID.Hex()), nil)
				So(err, ShouldBeNil)

				Convey("When not signed in", func() {
					res, err := http.DefaultClient.Do(req)
					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 401)

					dec := json.NewDecoder(res.Body)
					var e errors.Error
					So(dec.Decode(&e), ShouldBeNil)
					So(e.Code, ShouldEqual, 401)
				})

				Convey("When signed in", func() {
					user, _, err := tasks.CreateUser(&tasks.CreateUserRequest{
						Name:  "Test User",
						Email: "test@test.com",
					})
					So(err, ShouldBeNil)
					So(user, ShouldNotBeNil)

					token, _, err := tasks.CreateToken(&tasks.CreateTokenRequest{
						UserID: user.ID,
					})
					So(err, ShouldBeNil)

					req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

					Convey("Without admin permissions", func() {
						res, err := http.DefaultClient.Do(req)
						defer res.Body.Close()
						So(err, ShouldBeNil)

						So(res.StatusCode, ShouldEqual, 403)
					})

					Convey("With admin permissions", func() {
						_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
							UserID:      user.ID,
							Permissions: []string{"admin"},
						})
						So(err, ShouldBeNil)

						res, err := http.DefaultClient.Do(req)
						defer res.Body.Close()
						So(err, ShouldBeNil)

						So(res.StatusCode, ShouldEqual, 200)

						var e models.AuditLog
						dec := json.NewDecoder(res.Body)
						So(dec.Decode(&e), ShouldBeNil)
						So(e.ID, ShouldEqual, entry.ID)
						So(e.Timestamp.Unix(), ShouldEqual, entry.Timestamp.Unix())
						So(e.Token, ShouldEqual, entry.Token)
					})
				})
			})
		})

	})
}
