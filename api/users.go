package api

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
)

func init() {
	Router().
		Path("/v1/users").
		Methods("GET").
		Handler(girder.NewHandler(getUsers).
			RequireAuthentication(getUser).
			RequirePermission("admin/users").
			LogRequests()).
		Name("GET /users")
	Router().
		Path("/v1/users").
		Methods("POST").
		Handler(girder.NewHandler(createUser).
			RequireAuthentication(getUser).
			RequirePermission("admin/users").
			LogRequests()).
		Name("POST /users")

	Router().
		Path("/v1/user").
		Methods("GET").
		Handler(girder.NewHandler(getUserCurrent).
			RequireAuthentication(getUser).
			LogRequests()).
		Name("GET /user")

	Router().
		Path("/v1/user/{user}").
		Methods("GET").
		Handler(girder.NewHandler(getUserByID).
			RequireAuthentication(getUser).
			RequirePermission("admin/users").
			LogRequests()).
		Name("GET /user/{user}")

	Router().
		Path("/v1/user/{user}/permissions").
		Methods("POST").
		Handler(girder.NewHandler(addUserPermissions).
			RequireAuthentication(getUser).
			RequirePermission("admin/users").
			LogRequests()).
		Name("POST /user/{user}/permissions")
	Router().
		Path("/v1/user/{user}/permissions").
		Methods("PUT").
		Handler(girder.NewHandler(setUserPermissions).
			RequireAuthentication(getUser).
			RequirePermission("admin/users").
			LogRequests()).
		Name("PUT /user/{user}/permissions")
	Router().
		Path("/v1/user/{user}/permissions").
		Methods("DELETE").
		Handler(girder.NewHandler(removeUserPermissions).
			RequireAuthentication(getUser).
			RequirePermission("admin/users").
			LogRequests()).
		Name("PUT /user/{user}/permissions")
}

func getUserByID(c *girder.Context) (interface{}, error) {
	req := tasks.GetUserRequest{
		ID: c.Vars["user"],
	}

	user, err := tasks.GetUser(&req)
	if err != nil {
		return nil, errors.From(err)
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
		return nil, errors.From(err)
	}

	user, audit, err := tasks.CreateUser(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Token:   c.GetAuthToken().Value,
		User:    c.User.(*models.User).Summary(),
		Type:    "user.create",
		Context: audit,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return user, nil
}

func addUserPermissions(c *girder.Context) (interface{}, error) {
	req := tasks.AddPermissionsRequest{}

	if err := c.ReadBody(&req); err != nil {
		return nil, errors.From(err)
	}

	req.UserID = c.Vars["user"]

	user, audit, err := tasks.AddPermissions(&req)
	if err != nil {
		return nil, err
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Token:   c.GetAuthToken().Value,
		User:    c.User.(*models.User).Summary(),
		Type:    "user.permissions.add",
		Context: audit,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return user, nil
}

func setUserPermissions(c *girder.Context) (interface{}, error) {
	req := tasks.SetPermissionsRequest{}

	if err := c.ReadBody(&req); err != nil {
		return nil, errors.From(err)
	}

	req.UserID = c.Vars["user"]

	user, audit, err := tasks.SetPermissions(&req)
	if err != nil {
		return nil, err
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Token:   c.GetAuthToken().Value,
		User:    c.User.(*models.User).Summary(),
		Type:    "user.permissions.update",
		Context: audit,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return user, nil
}

func removeUserPermissions(c *girder.Context) (interface{}, error) {
	req := tasks.RemovePermissionsRequest{}

	if err := c.ReadBody(&req); err != nil {
		return nil, errors.From(err)
	}

	req.UserID = c.Vars["user"]

	user, audit, err := tasks.RemovePermissions(&req)
	if err != nil {
		return nil, err
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Token:   c.GetAuthToken().Value,
		User:    c.User.(*models.User).Summary(),
		Type:    "user.permissions.remove",
		Context: audit,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return user, nil
}
