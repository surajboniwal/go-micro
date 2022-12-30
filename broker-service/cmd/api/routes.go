package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Status:  true,
			Message: "Broker service for go-micro",
		}

		w.Header().Set("Content-Type", "application/json")

		res, _ := json.MarshalIndent(response, "", "  ")

		w.Write(res)
	})

	return mux
}
