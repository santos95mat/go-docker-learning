package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-docker-learning/internal/database"
)

func init() {
	database.Connect()
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
