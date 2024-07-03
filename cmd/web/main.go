package main

import (
	"log"
	"net/http"

	"github.com/srisudarshanrg/HackhiveProject/pkg/config"
	"github.com/srisudarshanrg/HackhiveProject/pkg/handlers"
	"github.com/srisudarshanrg/HackhiveProject/pkg/render"
)

const portNumber = ":5050"

var app config.AppConfig

func main() {
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Println("Could not create template cache.")
	}

	app.TemplateCache = templateCache

	render.SetAppConfig(&app)

	repo := handlers.SetAppConfigHandler(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal("Could not serve application", err)

}
