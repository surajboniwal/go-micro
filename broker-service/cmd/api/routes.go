package main

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world")
	})

	return mux
}
