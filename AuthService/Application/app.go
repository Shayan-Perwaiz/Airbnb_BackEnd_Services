package app

import (
	"fmt"
	"net/http"
)

type Config struct {
	Address string
}

type Application struct {
	Config Config
}


func NewConfig(address string) Config{
	cfg := Config{
		Address: address,
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
	}
	fmt.Println("Starting server on",app.Config.Address)
    return server.ListenAndServe()
}

