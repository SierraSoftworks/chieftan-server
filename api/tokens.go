package api

import (
	"github.com/SierraSoftworks/chieftan-server/tasks"
	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
)

func getTokens(c *girder.Context) (interface{}, error) {
	req := tasks.GetUserTokensRequest{
		ID: c.Vars["user"],
	}

	tokens, err := tasks.GetUserTokens(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return tokens, nil
}

func createToken(c *girder.Context) (interface{}, error) {
	req := tasks.CreateTokenRequest{
		UserID: c.Vars["user"],
	}

	token, err := tasks.CreateToken(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return struct {
		Token string `json:"token"`
	}{token}, nil
}

func revokeToken(c *girder.Context) (interface{}, error) {
	req := tasks.RemoveTokenRequest{
		Token: c.Vars["token"],
	}

	err := tasks.RemoveToken(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return nil, nil
}

func revokeAllTokens(c *girder.Context) (interface{}, error) {
	req := tasks.RemoveAllTokensRequest{}

	err := tasks.RemoveAllTokens(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return nil, nil
}

func revokeAllTokensForUser(c *girder.Context) (interface{}, error) {
	req := tasks.RemoveAllTokensRequest{
		UserID: c.Vars["user"],
	}

	err := tasks.RemoveAllTokens(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	return nil, nil
}

func init() {
	Router().Path("/v1/tokens").Methods("DELETE").Handler(girder.NewHandler(revokeAllTokens).RequirePermission("admin"))
	Router().Path("/v1/token/{token}").Methods("DELETE").Handler(girder.NewHandler(revokeToken).RequirePermission("admin/users"))

	Router().Path("/v1/user/{user}/tokens").Methods("GET").Handler(girder.NewHandler(getTokens).RequirePermission("admin/users"))
	Router().Path("/v1/user/{user}/tokens").Methods("POST").Handler(girder.NewHandler(createToken).RequirePermission("admin/users"))
	Router().Path("/v1/user/{user}/tokens").Methods("DELETE").Handler(girder.NewHandler(revokeAllTokensForUser).RequirePermission("admin/users"))
}
