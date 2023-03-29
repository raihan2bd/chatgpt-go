package middlewares

import (
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/raihan2bd/chatgpt-go/config"
)

var app *config.Application

func NewMiddlewares(a *config.Application) {
	app = a
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
