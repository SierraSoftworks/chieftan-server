package api

import (
	"github.com/SierraSoftworks/chieftan-server/tasks"

	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

// Router returns the registered router for the API
func Router() *mux.Router {
	return router
}

func getUser(token *girder.AuthorizationToken) (girder.User, error) {
	if token.Type != "Token" {
		return nil, errors.Unauthorized()
	}

	return tasks.GetUser(&tasks.GetUserRequest{
		Token: token.Value,
	})
}
