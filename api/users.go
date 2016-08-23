package api

import (
	"github.com/SierraSoftworks/chieftan-server/tasks"
	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
)

func getUserByID(c *girder.Context) (interface{}, error) {
	req := tasks.GetUserRequest{
		ID: c.Vars["user"],
	}

	user, err := tasks.GetUser(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	if user == nil {
		return nil, errors.NotFound()
	}

	return user, nil
}

func getUserCurrent(c *girder.Context) (interface{}, error) {
	req := tasks.GetUserRequest{
		ID: c.User.GetID(),
	}

	user, err := tasks.GetUser(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	if user == nil {
		return nil, errors.NotFound()
	}

	return user, nil
}

func getUsers(c *girder.Context) (interface{}, error) {
	req := tasks.GetUsersRequest{}

	users, err := tasks.GetUsers(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return users, nil
}

func createUser(c *girder.Context) (interface{}, error) {
	req := tasks.CreateUserRequest{}

	if err := c.ReadBody(&req); err != nil {
		return nil, err
	}

	user, err := tasks.CreateUser(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return user, nil
}

func init() {
	Router().Path("/v1/users").Methods("GET").Handler(girder.NewHandler(getUsers).LogRequests().RequireAuthentication(getUser).RequirePermission("admin/users"))
	Router().Path("/v1/users").Methods("POST").Handler(girder.NewHandler(createUser).RequireAuthentication(getUser).RequirePermission("admin/users"))

	Router().Path("/v1/user").Methods("GET").Handler(girder.NewHandler(getUserCurrent).RequireAuthentication(getUser))
	Router().Path("/v1/user/{user}").Methods("GET").Handler(girder.NewHandler(getUserByID).RequireAuthentication(getUser).RequirePermission("admin/users"))
}
