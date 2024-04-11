package handler

import (
	"net/http"
)

// A HealthzHandler implements health check endpoint.
type PanicHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewPanicHandler() *PanicHandler {
	return &PanicHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *PanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("Intentional panic for demonstration")
}
