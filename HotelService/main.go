package main

import (
	app "GoHotelService/apps"
	config "GoHotelService/configs/env"
	"fmt"
)


func main() {
	fmt.Println("this is Hotel Service")
	config.LoadEnv()
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)
	app.Run()
}