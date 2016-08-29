package api

import (
	"github.com/SierraSoftworks/chieftan-server/tasks"

	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func init() {
	router.NotFoundHandler = girder.NewHandler(func(c *girder.Context) (interface{}, error) {
		return nil, errors.NotFound()
	})
}

// Router returns the registered router for the API
func Router() *mux.Router {
	return router
}

func getUser(token *girder.AuthorizationToken) (girder.User, error) {
	if token.Type != "Token" {
		return nil, errors.Unauthorized()
	}

	user, err := tasks.GetUser(&tasks.GetUserRequest{
		Token: token.Value,
	})

	if err != nil {
		e := errors.From(err)
		if e.Code == 404 {
			return nil, errors.Unauthorized()
		}

		return nil, e
	}

	return user, nil
}
