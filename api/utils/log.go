package utils

import (
	log "github.com/Sirupsen/logrus"
)

// LogRequests adds debug logging on all requests to simplify diagnosing problems
func (h *Handler) LogRequests() *Handler {
	return h.RegisterPreprocessors(func(c *Context) *Error {
		log.WithFields(log.Fields{
			"headers": c.Request.Header,
			"vars":    c.Vars,
		}).Infof("%s %s", c.Request.Method, c.Request.URL.String())

		return nil
	})
}
