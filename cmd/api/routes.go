package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewMux()
	r.Get("/create", app.createGame)
	r.Get("/join", app.joinGame)
	return r
}
