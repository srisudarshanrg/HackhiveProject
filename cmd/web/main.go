package main

import (
	"log"
	"net/http"

	"github.com/srisudarshanrg/HackhiveProject/pkg/config"
	"github.com/srisudarshanrg/HackhiveProject/pkg/driver"
	"github.com/srisudarshanrg/HackhiveProject/pkg/handlers"
	"github.com/srisudarshanrg/HackhiveProject/pkg/render"
)

const portNumber = ":5000"

var app config.AppConfig

func main() {
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Println("Could not create template cache.")
	}

	app.TemplateCache = templateCache

	repo := handlers.SetAppConfigHandler(&app)
	handlers.NewHandlers(repo)
	render.SetAppConfig(&app)

	// create database
	db, err := driver.CreateDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	handlers.DatabaseAccess(db)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal("Could not serve application", err)

}
