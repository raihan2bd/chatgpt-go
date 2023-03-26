package config

import (
	"html/template"
	"log"
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
}
