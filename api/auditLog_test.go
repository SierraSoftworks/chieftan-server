package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestGetAuditLogEntry(c *C) {
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

	c.Assert(err, IsNil)

	url := fmt.Sprintf("%s%s%s", ts.URL, "/v1/audit/", entry.ID.Hex())

	req, err := http.NewRequest("GET", url, nil)
	c.Assert(err, IsNil)

	client := &http.Client{}

	res, err := client.Do(req)
	defer res.Body.Close()
	c.Assert(err, IsNil)
	c.Assert(res.StatusCode, Equals, 401)

	user, _, err := tasks.CreateUser(&tasks.CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})
	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	token, _, err := tasks.CreateToken(&tasks.CreateTokenRequest{
		UserID: user.ID,
	})
	c.Assert(err, IsNil)
	c.Assert(token, Not(Equals), "")

	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))

	res, err = client.Do(req)
	defer res.Body.Close()
	c.Assert(err, IsNil)
	c.Assert(res.StatusCode, Equals, 403)

	_, err = tasks.SetPermissions(&tasks.SetPermissionsRequest{
		UserID:      user.ID,
		Permissions: []string{"admin"},
	})
	c.Assert(err, IsNil)

	res, err = client.Do(req)
	defer res.Body.Close()
	c.Assert(err, IsNil)
	c.Check(res.StatusCode, Equals, 200)

	var e models.AuditLog
	dec := json.NewDecoder(res.Body)
	c.Assert(dec.Decode(&e), IsNil)
	c.Check(e.ID, Equals, entry.ID)
	c.Check(e.Timestamp.Unix(), Equals, entry.Timestamp.Unix())
	c.Check(e.Token, Equals, entry.Token)
}
