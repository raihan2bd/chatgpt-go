package handlers

import (
	"net/http"

	"github.com/raihan2bd/chatgpt-go/config"
	"github.com/raihan2bd/chatgpt-go/models"
	"github.com/raihan2bd/chatgpt-go/render"
	"golang.org/x/crypto/bcrypt"
)

var app *config.Application

func NewHandlers(a *config.Application) {
	app = a
}

// HomeHandler displays home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	user, err := app.DB.GetUserByID(1)
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}
	app.InfoLog.Println(user.FirstName, user.LastName)
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

// PostSignupHandler create new account and save user credentials to the database
func PostSignupHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	user := models.User{
		FirstName:   "Jakir",
		LastName:    "hossen",
		Email:       "me@r.com",
		Password:    "123456",
		AccessLevel: 1,
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		app.ErrorLog.Println(err)
	}

	err = app.DB.AddNewUser(user, string(hash))
	if err != nil {
		app.ErrorLog.Println(err)
	}

	http.Redirect(w, r, "/chatgpt", http.StatusSeeOther)
}

// SignupHandler displays signup page
func ChatGptHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "chatgpt.page.html", &models.TemplateData{})
}
