package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
func (app *application) routes() *chi.Mux {
	r := chi.NewMux()
	r.Get("/ping", ping)
	r.Get("/v1/create", app.createGameHandler) //
	r.Get("/v1/join", app.joinGameHandler)
	r.Get("/v1/games", app.listGameHandler)
	r.Post("/v1/signup", app.signupHandler)
	r.Post("/v1/signin", app.signinHandler)
	return r
}
