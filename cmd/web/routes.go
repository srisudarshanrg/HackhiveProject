package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/srisudarshanrg/HackhiveProject/pkg/config"
	"github.com/srisudarshanrg/HackhiveProject/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/login", handlers.Repo.Login)
	mux.Post("/login", handlers.Repo.PostLogin)

	mux.Get("/signup", handlers.Repo.SignUp)
	mux.Post("/signup", handlers.Repo.PostSignUp)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
