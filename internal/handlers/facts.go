package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-docker-learning/internal/database"
	"github.com/santos95mat/go-docker-learning/internal/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.DB.Find(&facts)

	return c.Status(fiber.StatusFound).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := models.Fact{}

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.DB.Create(&fact)

	return c.Status(fiber.StatusCreated).JSON(fact)
}
