package handlers

import (
	"net/http"

	"github.com/raihan2bd/chatgpt-go/models"
	"github.com/raihan2bd/chatgpt-go/render"
)

// HomeHandler displays home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

// LoginHandler displays login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "login.page.html", &models.TemplateData{})
}

// SignupHandler displays signup page
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "signup.page.html", &models.TemplateData{})
}

// SignupHandler displays signup page
func ChatGptHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "chatgpt.page.html", &models.TemplateData{})
}
