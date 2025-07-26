package app

import (
	dbconfig "GoReview/configs/db"
	config "GoReview/configs/env"
	"GoReview/repository/db"
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
	port := config.GetString("PORT", ":8082")
	return Config{
		Address: port,
	}
}

func NewApplication(_cfg Config) *Application{
	return &Application{
		Config: _cfg,
	}
}

func(app *Application) Run(){
	dbConnector, err := dbconfig.SetUpDB()
	if err != nil{
		fmt.Println("Error", err)
	}
	reviewRepo := db.NewReviewRepositoryImpl(dbConnector)

	server := &http.Server{
		Addr : app.Config.Address,
	    ReadTimeout:    10 * time.Second,
	    WriteTimeout:   10 * time.Second,
	}
	fmt.Println("Server is running on port", server.Addr)
	server.ListenAndServe()
}
