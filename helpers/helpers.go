package helpers

import (
	"net/http"

	"github.com/raihan2bd/chatgpt-go/config"
)

var app *config.Application

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.Application) {
	app = a
}

func IsAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "user_id")
	return exists
}
