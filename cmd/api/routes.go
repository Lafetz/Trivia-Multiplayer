package main

import "github.com/go-chi/chi/v5"

func (app *application) routes() *chi.Mux {
	r := chi.NewMux()
	r.Get("/v1/create", app.createGame)
	r.Get("/v1/join", app.joinGame)
	r.Get("/v1/signup", app.signupHandler)
	r.Get("/v1/signin", app.signinHandler)
	return r
}
