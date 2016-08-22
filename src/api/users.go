package api

import (
	"github.com/SierraSoftworks/chieftan-server/src/api/utils"
	"github.com/SierraSoftworks/chieftan-server/src/tasks"
)

func getUserByID(c *utils.Context) (interface{}, *utils.Error) {
	req := tasks.GetUserRequest{
		ID: c.Vars["user"],
	}

	user, err := tasks.GetUser(&req)
	if err != nil {
		return nil, utils.NewErrorFor(err)
	}

	if user == nil {
		return nil, utils.NotFound()
	}

	return user, nil
}

func getUserCurrent(c *utils.Context) (interface{}, *utils.Error) {
	req := tasks.GetUserRequest{
		ID: c.User.ID(),
	}

	user, err := tasks.GetUser(&req)
	if err != nil {
		return nil, utils.NewErrorFor(err)
	}

	if user == nil {
		return nil, utils.NotFound()
	}

	return user, nil
}

func getUsers(c *utils.Context) (interface{}, *utils.Error) {
	req := tasks.GetUsersRequest{}

	users, err := tasks.GetUsers(&req)
	if err != nil {
		return nil, utils.NewErrorFor(err)
	}

	return users, nil
}

func createUser(c *utils.Context) (interface{}, *utils.Error) {
	req := tasks.CreateUserRequest{}

	if err := c.ReadBody(&req); err != nil {
		return nil, err
	}

	user, err := tasks.CreateUser(&req)
	if err != nil {
		return nil, utils.NewErrorFor(err)
	}

	return user, nil
}

func init() {
	Router().Path("/v1/users").Methods("GET").Handler(utils.NewHandler(getUsers).LogRequests().RequirePermissions("admin/users"))
	Router().Path("/v1/users").Methods("POST").Handler(utils.NewHandler(createUser).RequirePermissions("admin/users"))

	Router().Path("/v1/user").Methods("GET").Handler(utils.NewHandler(getUserCurrent))
	Router().Path("/v1/user/{user}").Methods("GET").Handler(utils.NewHandler(getUserByID).RequirePermissions("admin/users"))
}
