package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	. "github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()

	healthzHandler := Recovery(Device(AccessLog(handler.NewHealthzHandler())))
	mux.Handle("/healthz", healthzHandler)

	todoService := service.NewTODOService(todoDB)
	todoHandler := Recovery(Auth(Device(AccessLog(handler.NewTODOHandler(todoService)))))
	mux.Handle("/todos", todoHandler)

	panicHandler := Recovery(Device(AccessLog(handler.NewPanicHandler())))
	mux.Handle("/do-panic", panicHandler)

	shutdownHandler := Recovery(Device(AccessLog(handler.NewShutdownHandler())))
	mux.Handle("/shutdown", shutdownHandler)

	return mux
}
