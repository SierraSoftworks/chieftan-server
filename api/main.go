package api

import (
	"github.com/SierraSoftworks/chieftan-server/api/utils"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

// Router returns the registered router for the API
func Router() *mux.Router {
	return router
}

func init() {
	utils.ActiveUserStore = &userStore{}
}
