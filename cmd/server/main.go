package main

import "github.com/Valery223/Question-Bank/internal/app"

func main() {
	app := app.NewApp()
	app.Run(":8080")
}
