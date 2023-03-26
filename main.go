package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/raihan2bd/chatgpt-go/config"
	"github.com/raihan2bd/chatgpt-go/routes"
)

var app config.Application
var port string

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("No environment variable found!")
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app.Config.Port = port
	app.InfoLog = infoLog
	app.ErrorLog = errorLog

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
