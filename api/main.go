package api

import (
	"github.com/SierraSoftworks/chieftan-server/tasks"

	"github.com/SierraSoftworks/girder"
	"github.com/SierraSoftworks/girder/errors"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func init() {
	router.NotFoundHandler = notFoundHandler
	router.StrictSlash(true)
}

var notFoundHandler = girder.NewHandler(func(c *girder.Context) (interface{}, error) {
	log.WithFields(log.Fields{
		"url":        c.Request.URL,
		"method":     c.Request.Method,
		"user-agent": c.Request.UserAgent(),
		"headers":    c.Request.Header,
	}).Info("Route Not Found")
	return nil, errors.NotFound()
})

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
