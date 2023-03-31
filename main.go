package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"github.com/raihan2bd/chatgpt-go/config"
	"github.com/raihan2bd/chatgpt-go/driver"
	"github.com/raihan2bd/chatgpt-go/handlers"
	"github.com/raihan2bd/chatgpt-go/helpers"
	"github.com/raihan2bd/chatgpt-go/middlewares"
	"github.com/raihan2bd/chatgpt-go/render"
	"github.com/raihan2bd/chatgpt-go/routes"
	"github.com/sashabaranov/go-openai"
)

var app config.Application
var port string
var production bool
var session *scs.SessionManager

func main() {
	// load environment variables file
	_ = godotenv.Load()

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

	// open ai secret
	openAIKey := os.Getenv("OPEN_AI_KEY")
	app.Config.OpenAIKey = openAIKey

	// initialization session
	gob.Register(map[string]int{})
	app.OpenAIClients = openai.NewClient(app.Config.OpenAIKey)
	session = scs.New()
	session.Store = postgresstore.New(conn)
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// create templates cache
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

	// share application data
	handlers.NewHandlers(&app)
	render.NewTemplates(&app)
	middlewares.NewMiddlewares(&app)
	helpers.NewHelpers(&app)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: routes.Routes(),
		// IdleTimeout:       30 * time.Second,
		// ReadTimeout:       10 * time.Second,
		// ReadHeaderTimeout: 5 * time.Second,
		// WriteTimeout:      5 * time.Second,
	}

	app.InfoLog.Printf("Starting HTTP server on port %s\n", app.Config.Port)
	err = srv.ListenAndServe()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Fatal(err)
	}
}
