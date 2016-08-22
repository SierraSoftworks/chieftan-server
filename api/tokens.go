package api

import (
	"github.com/SierraSoftworks/chieftan-server/api/utils"
	"github.com/SierraSoftworks/chieftan-server/tasks"
)

func getTokens(c *utils.Context) (interface{}, *utils.Error) {
	req := tasks.GetUserTokensRequest{
		ID: c.Vars["user"],
	}

	tokens, err := tasks.GetUserTokens(&req)
	if err != nil {
		return nil, utils.NewErrorFor(err)
	}

	return tokens, nil
}

func createToken(c *utils.Context) (interface{}, *utils.Error) {
	req := tasks.CreateTokenRequest{
		UserID: c.Vars["user"],
	}

	token, err := tasks.CreateToken(&req)
	if err != nil {
		return nil, utils.NewErrorFor(err)
	}

	return struct {
		Token string `json:"token"`
	}{token}, nil
}

func revokeToken(c *utils.Context) (interface{}, *utils.Error) {
	req := tasks.RemoveTokenRequest{
		Token: c.Vars["token"],
	}

	err := tasks.RemoveToken(&req)
	if err != nil {
		return nil, utils.NewErrorFor(err)
	}

	return nil, nil
}

func revokeAllTokens(c *utils.Context) (interface{}, *utils.Error) {
	req := tasks.RemoveAllTokensRequest{}

	err := tasks.RemoveAllTokens(&req)
	if err != nil {
		return nil, utils.NewErrorFor(err)
	}

	return nil, nil
}

func revokeAllTokensForUser(c *utils.Context) (interface{}, *utils.Error) {
	req := tasks.RemoveAllTokensRequest{
		UserID: c.Vars["user"],
	}

	err := tasks.RemoveAllTokens(&req)
	if err != nil {
		return nil, utils.NewErrorFor(err)
	}

	return nil, nil
}

func init() {
	Router().Path("/v1/tokens").Methods("DELETE").Handler(utils.NewHandler(revokeAllTokens).RequirePermissions("admin"))
	Router().Path("/v1/token/{token}").Methods("DELETE").Handler(utils.NewHandler(revokeToken).RequirePermissions("admin/users"))

	Router().Path("/v1/user/{user}/tokens").Methods("GET").Handler(utils.NewHandler(getTokens).RequirePermissions("admin/users"))
	Router().Path("/v1/user/{user}/tokens").Methods("POST").Handler(utils.NewHandler(createToken).RequirePermissions("admin/users"))
	Router().Path("/v1/user/{user}/tokens").Methods("DELETE").Handler(utils.NewHandler(revokeAllTokensForUser).RequirePermissions("admin/users"))
}
