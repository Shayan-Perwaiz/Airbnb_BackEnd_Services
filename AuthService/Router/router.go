package router

import (
	controller "GoAuth/Controllers"

	"github.com/go-chi/chi/v5"
)

func SetUpRouter() *chi.Mux{
	r := chi.NewRouter()
	r.Get("/ping", controller.PingHandler)
	return r
}