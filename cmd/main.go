package main

import (
	"log"

	"ai.zustack.backend/api"
	"ai.zustack.backend/internal/database"
)

func main() {
	database.ConnectDB()
	app := api.Setup()
	log.Fatal(app.Listen(":8080"))
}
