package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/raihan2bd/chatgpt-go/models"
)

type Config struct {
	Port   string
	DSN    string
	GPTKey string
}

type Application struct {
	Config        Config
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	TemplateCache map[string]*template.Template
	InProduction  bool
	DB            models.DBModel
	Session       *scs.SessionManager
}
