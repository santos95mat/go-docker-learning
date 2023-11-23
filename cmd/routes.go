package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-docker-learning/internal/handler"
	"github.com/santos95mat/go-docker-learning/internal/repository"
)

var (
	factRepository = repository.NewFactRepository()
	factHandler    = handler.NewFactHandler(factRepository)
)

func setupRoutes(app *fiber.App) {
	app.Post("/fact", factHandler.CreateFact)
	app.Get("/fact", factHandler.ListFacts)
	app.Get("/fact/:id", factHandler.ListFact)
}
