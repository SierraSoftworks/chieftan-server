package api

import (
	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/chieftan-server/tasks"
	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
)

func init() {
	Router().
		Path("/v1/tokens").
		Methods("DELETE").
		Handler(girder.NewHandler(revokeAllTokens).
			RequireAuthentication(getUser).
			RequirePermission("admin").
			LogRequests()).
		Name("DELETE /tokens")
	Router().
		Path("/v1/token/{token}").
		Methods("DELETE").
		Handler(girder.NewHandler(revokeToken).
			RequireAuthentication(getUser).
			RequirePermission("admin/users").
			LogRequests()).
		Name("DELETE /token/{token}")

	Router().
		Path("/v1/user/{user}/tokens").
		Methods("GET").
		Handler(girder.NewHandler(getTokens).
			RequireAuthentication(getUser).
			RequirePermission("admin/users").
			LogRequests()).
		Name("GET /user/{user}/tokens")
	Router().
		Path("/v1/user/{user}/tokens").
		Methods("POST").
		Handler(girder.NewHandler(createToken).
			RequireAuthentication(getUser).
			RequirePermission("admin/users").
			LogRequests()).
		Name("POST /user/{user}/tokens")
	Router().
		Path("/v1/user/{user}/tokens").
		Methods("DELETE").
		Handler(girder.NewHandler(revokeAllTokensForUser).
			RequireAuthentication(getUser).
			RequirePermission("admin/users").
			LogRequests()).
		Name("DELETE /user/{user}/tokens")

	Router().
		Path("/v1/user/{user}/token/{token}").
		Methods("DELETE").
		Handler(girder.NewHandler(revokeToken).
			RequireAuthentication(getUser).
			RequirePermission("admin/users").
			LogRequests()).
		Name("DELETE /user/{user}/token/{token}")

}

func getTokens(c *girder.Context) (interface{}, error) {
	req := tasks.GetTokensRequest{
		UserID: c.Vars["user"],
	}

	tokens, audit, err := tasks.GetTokens(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Token:   c.GetAuthToken().Value,
		User:    c.User.(*models.User).Summary(),
		Type:    "user.tokens.view",
		Context: audit,
	})

	if err != nil {
		return nil, errors.From(err)
	}

	return tokens, nil
}

func createToken(c *girder.Context) (interface{}, error) {
	req := tasks.CreateTokenRequest{
		UserID: c.Vars["user"],
	}

	token, audit, err := tasks.CreateToken(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Token:   c.GetAuthToken().Value,
		User:    c.User.(*models.User).Summary(),
		Type:    "user.tokens.create",
		Context: audit,
	})
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

	audit, err := tasks.RemoveToken(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Token:   c.GetAuthToken().Value,
		User:    c.User.(*models.User).Summary(),
		Type:    "user.tokens.revoke",
		Context: audit,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return nil, nil
}

func revokeAllTokens(c *girder.Context) (interface{}, error) {
	req := tasks.RemoveAllTokensRequest{}

	audit, err := tasks.RemoveAllTokens(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Token:   c.GetAuthToken().Value,
		User:    c.User.(*models.User).Summary(),
		Type:    "tokens.revoke",
		Context: audit,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return nil, nil
}

func revokeAllTokensForUser(c *girder.Context) (interface{}, error) {
	req := tasks.RemoveAllTokensRequest{
		UserID: c.Vars["user"],
	}

	audit, err := tasks.RemoveAllTokens(&req)
	if err != nil {
		return nil, errors.From(err)
	}

	_, err = tasks.CreateAuditLogEntry(&tasks.CreateAuditLogEntryRequest{
		Token:   c.GetAuthToken().Value,
		User:    c.User.(*models.User).Summary(),
		Type:    "user.tokens.revoke",
		Context: audit,
	})
	if err != nil {
		return nil, errors.From(err)
	}

	return nil, nil
}
