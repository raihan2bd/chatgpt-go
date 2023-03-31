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
	router.Use(middlewares.SessionLoad)

	router.Get("/", handlers.HomeHandler)

	router.Get("/login", handlers.LoginHandler)
	router.Post("/login", handlers.PostLoginHandler)
	router.Get("/signup", handlers.SignupHandler)
	router.Post("/signup", handlers.PostSignupHandler)
	router.Get("/logout", handlers.LogoutHandler)

	router.Route("/chatgpt", func(router chi.Router) {
		router.Use(middlewares.Auth)
		router.Get("/", handlers.ChatGptHandler)
		router.Post("/", handlers.PostChatGptHandler)
	})

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}
