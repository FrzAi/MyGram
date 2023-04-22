package main

import (
	"Final_Project/app"
	"Final_Project/config"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	err := config.InitPostgres()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
