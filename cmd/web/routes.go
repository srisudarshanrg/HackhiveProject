package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/srisudarshanrg/HackhiveProject/pkg/config"
	"github.com/srisudarshanrg/HackhiveProject/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Login)
	mux.Post("/", handlers.Repo.PostLogin)

	mux.Get("/signin", handlers.Repo.SignUp)
	mux.Post("/signin", handlers.Repo.PostSignUp)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
