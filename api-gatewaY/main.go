package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/golang-jwt/jwt/v5"
)

// TOKEM
var jwtSecret = []byte("dark0souls")

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		log.Printf("[%s] %s", c.Method(), c.OriginalURL())
		return c.Next()
	})

	// Rotas de autenticação — não exigem JWT
	app.All("/auth/*", func(c *fiber.Ctx) error {
		newPath := c.OriginalURL()[len("/auth"):]
		target := "http://127.0.0.1:3001" + newPath
		return proxy.Do(c, target)
	})

	// Middleware para verificar JWT nas rotas /api/*
	app.Use("/api/*", func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token não fornecido",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Método de assinatura inválido")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido ou expirado",
			})
		}

		return c.Next()
	})

	app.All("/api/*", func(c *fiber.Ctx) error {
		newPath := c.OriginalURL()[len("/api"):]
		target := "http://127.0.0.1:3002" + newPath
		return proxy.Do(c, target)
	})

	log.Fatal(app.Listen(":3000"))
}
