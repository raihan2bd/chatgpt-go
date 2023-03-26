package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raihan2bd/chatgpt-go/handlers"
)

func Routes() http.Handler {
	// Initialize
	router := chi.NewRouter()

	router.Get("/", handlers.HomeHandler)

	return router
}
