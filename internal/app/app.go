package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-docker-learning/internal/app/routes"
)

func Run() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	routes.FactRoutes(app)
}
