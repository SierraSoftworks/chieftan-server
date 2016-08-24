package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/SierraSoftworks/chieftan-server/tasks"
	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestGetUsers(c *C) {
	ts := httptest.NewServer(Router())
	defer ts.Close()
	url := fmt.Sprintf("%s%s", ts.URL, "/v1/users")

	user, _, err := tasks.CreateUser(&tasks.CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})

	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	req, err := http.NewRequest("GET", url, nil)
	c.Assert(err, IsNil)

	client := &http.Client{}

	res, err := client.Do(req)
	defer res.Body.Close()
	c.Assert(err, IsNil)
	c.Assert(res.StatusCode, Equals, 401)

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
		Permissions: []string{"admin/users"},
	})
	c.Assert(err, IsNil)

	res, err = client.Do(req)
	defer res.Body.Close()
	c.Assert(err, IsNil)
	c.Check(res.StatusCode, Equals, 200)
}

func (s *TestSuite) TestGetUser(c *C) {
	ts := httptest.NewServer(Router())
	defer ts.Close()

	user, _, err := tasks.CreateUser(&tasks.CreateUserRequest{
		Name:  "Test User",
		Email: "test@test.com",
	})

	url := fmt.Sprintf("%s%s%s", ts.URL, "/v1/user/", user.ID)

	c.Assert(err, IsNil)
	c.Assert(user, NotNil)

	req, err := http.NewRequest("GET", url, nil)
	c.Assert(err, IsNil)

	client := &http.Client{}

	res, err := client.Do(req)
	defer res.Body.Close()
	c.Assert(err, IsNil)
	c.Assert(res.StatusCode, Equals, 401)

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
		Permissions: []string{"admin/users"},
	})
	c.Assert(err, IsNil)

	res, err = client.Do(req)
	defer res.Body.Close()
	c.Assert(err, IsNil)
	c.Assert(res.StatusCode, Equals, 200)
}
