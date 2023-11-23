package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-docker-learning/internal/database"
	"github.com/santos95mat/go-docker-learning/internal/model"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []model.Fact{}

	database.DB.Find(&facts)

	return c.Status(fiber.StatusFound).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := model.Fact{}

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Create(&fact)

	return c.Status(fiber.StatusCreated).JSON(fact)
}
