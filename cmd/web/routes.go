package web

import (
	"net/http"
	"ws/internal/handlers"

	"github.com/bmizerany/pat"
)

// routes defines the application routes
func Routes() http.Handler {
	mux := pat.New()

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	return mux
}
