package app

import (
	config "GoAuth/Config/env"
	controller "GoAuth/Controllers"
	router "GoAuth/Router"
	db "GoAuth/db/repositories"
	"GoAuth/services"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Address string
}

type Application struct {
	Config Config
	Store db.Storage
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

	ur := db.NewUserRepository()
	us := services.NewUserService(ur)
	uc := controller.NewUserController(us)
	uRouter := router.NewUserRouter(uc)


	server := &http.Server{
		Addr : app.Config.Address,
		Handler: router.SetUpRouter(uRouter),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on",app.Config.Address)
    return server.ListenAndServe()
}

