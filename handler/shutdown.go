package handler

import (
	"context"
	"fmt"
	"net/http"
)

// A HealthzHandler implements health check endpoint.
type ShutdownHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewShutdownHandler() *ShutdownHandler {
	return &ShutdownHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *ShutdownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST method required", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Shutting down the server..."))

	if srv := r.Context().Value(http.ServerContextKey); srv != nil {
		if server, ok := srv.(*http.Server); ok {
			fmt.Println("Server shutted down")
			server.Shutdown(context.Background())
		}
	}
}
