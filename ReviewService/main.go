package main

import (
	app "GoReview/application"
	config "GoReview/configs/env"
	"fmt"
)

func main() {
	fmt.Println("This is Shayan")
	config.LoadEnv()
	cfg := app.NewConfig()
	app := app.NewApplication(cfg)
	app.Run()
	
}