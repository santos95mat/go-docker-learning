package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-docker-learning/internal/handler"
)

func setupRoutes(app *fiber.App) {
	app.Get("/fact", handler.ListFacts)
	app.Post("/fact", handler.CreateFact)
}
