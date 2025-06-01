package main

import (
	_ "auth-service/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func setupRoutes(app *fiber.App) {
	app.Post("/register", register)
	app.Post("/login", login)

	//Pega a rota apenas direta
	app.Get("/swagger/*", swagger.HandlerDefault)
}
