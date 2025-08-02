package app

import (
	config "GoHotelService/configs/env"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Address string
}

type Application struct {
	Config Config
}

func NewConfig() Config{
	port := config.GetString("PORT_NUMBER", ":8003")
	return Config{
		Address: port,
	}
}

func NewApplication(_cfg Config) *Application{
	return &Application{
		Config: _cfg,
	}
}

func(app *Application) Run() error{
	server := &http.Server{
	Addr : app.Config.Address,
	ReadTimeout : 10 * time.Second,
	WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on",app.Config.Address)
	return server.ListenAndServe()
}