package main

import (
	_ "product-service/docs"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"
)

func setupRoutes(app *fiber.App) {
	app.Get("/products", getProducts)
	app.Get("/products/:id", getProduct)
	app.Post("/products", createProduct)
	app.Put("/products/:id", updateProduct)
	app.Delete("/products/:id", deleteProduct)

	//Pega a rota apenas direta
	app.Get("/swagger/*", swagger.HandlerDefault)
}
