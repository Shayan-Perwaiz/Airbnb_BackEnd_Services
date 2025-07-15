package main

import (
	config "GoAuth/Config/env"
	app "GoAuth/app"
)

func main() {
	config.Load()
	config := app.NewConfig()
	app := app.NewApplication(config)
	app.Run()
}