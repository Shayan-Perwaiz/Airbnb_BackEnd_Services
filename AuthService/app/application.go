package app

import (
	config "GoAuth/Config/env"
	router "GoAuth/Router"
	"fmt"
	"net/http"
)

type Config struct {
	Address string
}

type Application struct {
	Config Config
}


func NewConfig() Config{
	port := config.GetString("PORT", ":8080")
	cfg := Config{
		Address: port,
	}
	return cfg
}

func NewApplication(cfg Config) *Application{
	app := Application{
		Config: cfg,
	}
	return &app
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr : app.Config.Address,
		Handler: router.SetUpRouter(),
	}
	fmt.Println("Starting server on",app.Config.Address)
    return server.ListenAndServe()
}

