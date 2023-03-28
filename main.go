package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/raihan2bd/chatgpt-go/config"
	"github.com/raihan2bd/chatgpt-go/driver"
	"github.com/raihan2bd/chatgpt-go/handlers"
	"github.com/raihan2bd/chatgpt-go/middlewares"
	"github.com/raihan2bd/chatgpt-go/render"
	"github.com/raihan2bd/chatgpt-go/routes"
)

var app config.Application
var port string
var production bool

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("No environment variable found!")
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	production = os.Getenv("PRODUCTION") == "true"

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// connect with postgresql database
	dsn := os.Getenv("DATABASE_URI")
	conn, err := driver.ConnectSQL(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create templates cache")
		return
	}

	app.Config.Port = port
	app.InfoLog = infoLog
	app.ErrorLog = errorLog
	app.InProduction = production
	app.TemplateCache = templateCache
	app.DB.DB = conn

	// share data
	handlers.NewHandlers(&app)
	render.NewTemplates(&app)
	middlewares.NewMiddlewares(&app)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", app.Config.Port),
		Handler:           routes.Routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.InfoLog.Printf("Starting HTTP server on port %s\n", app.Config.Port)
	err = srv.ListenAndServe()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Fatal(err)
	}
}
