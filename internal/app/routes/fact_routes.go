package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-docker-learning/internal/handler"
	"github.com/santos95mat/go-docker-learning/internal/repository"
)

var (
	factRepository = repository.NewFactRepository()
	factHandler    = handler.NewFactHandler(factRepository)
)

func FactRoutes(app *fiber.App) {
	fact := app.Group("/fact")

	fact.Post("/", factHandler.CreateFact)
	fact.Get("/", factHandler.ListFacts)
	fact.Get("/:id", factHandler.ListFact)
}
