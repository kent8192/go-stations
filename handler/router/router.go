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

	healthzHandler := AccessLog(Device(Recovery(handler.NewHealthzHandler())))
	mux.Handle("/healthz", healthzHandler)

	todoService := service.NewTODOService(todoDB)
	todoHandler := Device(Recovery(handler.NewTODOHandler(todoService)))
	mux.Handle("/todos", todoHandler)

	panicHandler := Device(Recovery(handler.NewPanicHandler()))
	mux.Handle("/do-panic", panicHandler)
	return mux
}
