package middlewares

import (
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/raihan2bd/chatgpt-go/config"
	"github.com/raihan2bd/chatgpt-go/helpers"
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

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuthenticated(r) {
			app.Session.Put(r.Context(), "error", "Log in first!")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

		next.ServeHTTP(w, r)
	})
}
