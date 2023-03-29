package models

import "github.com/raihan2bd/chatgpt-go/forms"

// TemplateData sent data from handler to templates

type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Warning         string
	Flash           string
	Error           string
	IsAuthenticated int
	Form            *forms.Form
	UserRole        int
}
