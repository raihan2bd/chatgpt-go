package helpers

import (
	"encoding/json"
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

// writeJSON writes aribtrary data out as JSON
func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)

	return nil
}
