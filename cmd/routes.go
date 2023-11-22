package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-docker-learning/internal/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/fact", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)
}
