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

	mux.Get("/forgotpassword", handlers.Repo.ForgotPassword)
	mux.Post("/forgotpassword", handlers.Repo.PostForgotPassword)

	mux.Get("/otpconfirm", handlers.Repo.ConfirmOTP)
	mux.Post("/otpconfirm", handlers.Repo.PostConfirmOTP)

	mux.Get("/resetpassword", handlers.Repo.ResetPassword)
	mux.Post("/resetpassword", handlers.Repo.PostResetPassword)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
