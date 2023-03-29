package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raihan2bd/chatgpt-go/handlers"
	"github.com/raihan2bd/chatgpt-go/middlewares"
)

func Routes() http.Handler {
	// Initialize the router
	router := chi.NewRouter()
	router.Use(middlewares.NoSurf)

	router.Get("/", handlers.HomeHandler)
	router.Get("/login", handlers.LoginHandler)
	router.Get("/signup", handlers.SignupHandler)
	router.Post("/signup", handlers.PostSignupHandler)
	router.Route("/chatgpt", func(router chi.Router) {
		router.Get("/", handlers.ChatGptHandler)
	})

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}
