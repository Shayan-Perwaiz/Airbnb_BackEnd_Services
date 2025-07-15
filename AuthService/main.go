package main

import app "GoAuth/Application"

func main() {
	config := app.NewConfig(":3001")
	app := app.NewApplication(config)
	app.Run()
}