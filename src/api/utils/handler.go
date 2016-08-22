package utils

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// HandlerFunc represents a function that handles an API request and returns a result
type HandlerFunc func(c *Context) (interface{}, *Error)

// Handler represents an API request handler
type Handler struct {
	HandleFunc HandlerFunc

	Preprocessors []Preprocessor
}

// NewHandler creates a new HTTP compatible handler for the given API HandlerFunc
func NewHandler(handle HandlerFunc) *Handler {
	return &Handler{
		HandleFunc:    handle,
		Preprocessors: []Preprocessor{},
	}
}

func (h *Handler) RegisterPreprocessors(preprocessors ...Preprocessor) *Handler {
	h.Preprocessors = append(h.Preprocessors, preprocessors...)
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{
		Request:         r,
		ResponseHeaders: w.Header(),
		Vars:            mux.Vars(r),
		StatusCode:      200,

		response: w,
	}

	authToken := c.GetAuthToken()
	if authToken != nil && ActiveUserStore != nil {
		user, err := ActiveUserStore.GetUser(authToken)
		if err != nil {
			w.WriteHeader(err.Code)
			if err := writeJSON(err, c); err != nil {
				log.Error("Failed to encode error to JSON", err)
				return
			}
		}

		c.User = user
	}

	for _, preprocessor := range h.Preprocessors {
		if err := preprocessor(c); err != nil {
			w.WriteHeader(err.Code)
			if err := writeJSON(err, c); err != nil {
				log.Error("Failed to encode error to JSON", err)
			}

			return
		}
	}

	res, err := h.HandleFunc(c)

	if res == nil && err == nil {
		err = NewError(404, "Not Found", "We could not find the entity you were looking for. Please check your request and try again.")
	}

	if err != nil {
		w.WriteHeader(err.Code)
		if err := writeJSON(err, c); err != nil {
			log.Error("Failed to encode error to JSON", err)
		}

		return
	}

	w.WriteHeader(c.StatusCode)
	if err := writeJSON(res, c); err != nil {
		log.WithField("response", res).Error("Failed to encode response to JSON", err)
	}
}
