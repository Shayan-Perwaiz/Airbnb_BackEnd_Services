package router

import "github.com/go-chi/chi/v5"

type Router interface{
	RegisterRouter(r chi.Router)
}

func SetUpRouter(ReviewRouter Router) *chi.Mux{
	chiRouter := chi.NewRouter()
	ReviewRouter.RegisterRouter(chiRouter)
	return chiRouter
}