package router

import (
	controller "GoAuth/Controllers"

	"github.com/go-chi/chi/v5"
)
type Router interface{
	Register(r chi.Router)
}
func SetUpRouter(UserRouter Router) *chi.Mux{
	chiRouter := chi.NewRouter()
	chiRouter.Get("/ping", controller.PingHandler)
	UserRouter.Register(chiRouter)
	return chiRouter
}