package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewMux()
	r.Get("/socket", app.socketHandler)

	return r
}
