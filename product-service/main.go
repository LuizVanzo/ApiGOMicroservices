package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// @title Auth Service API
// @version 1.0
// @description API de autenticação com JWT
// @host localhost:3002
// @BasePath
func main() {
	InitDatabase()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3002"))
}
