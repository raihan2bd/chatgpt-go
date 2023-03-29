package handlers

import (
	"net/http"

	"github.com/raihan2bd/chatgpt-go/config"
	"github.com/raihan2bd/chatgpt-go/forms"
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
	render.RenderTemplate(w, r, "signup.page.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostSignupHandler create new account and save user credentials to the database
func PostSignupHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	user := models.User{
		FirstName:   r.Form.Get("first_name"),
		LastName:    r.Form.Get("last_name"),
		Email:       r.Form.Get("email"),
		Password:    r.Form.Get("password"),
		AccessLevel: 1,
	}

	// validate the form
	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email", "password")
	form.MinLength("first_name", 3)
	form.MaxLength("first_name", 30)
	form.MaxLength("last_name", 30)
	form.IsEmail("email")
	form.MinLength("password", 6)
	form.MaxLength("password", 11)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["signup_form"] = user
		render.RenderTemplate(w, r, "signup.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
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
