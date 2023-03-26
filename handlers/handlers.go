package handlers

import (
	"net/http"

	"github.com/raihan2bd/chatgpt-go/models"
	"github.com/raihan2bd/chatgpt-go/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}
