package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-docker-learning/internal/dto"
	"github.com/santos95mat/go-docker-learning/internal/interfaces"
)

type factHandler struct {
	factRepository interfaces.FactRepository
}

func NewFactHandler(repo interfaces.FactRepository) *factHandler {
	return &factHandler{factRepository: repo}
}

func (f *factHandler) ListFacts(c *fiber.Ctx) error {
	facts := f.factRepository.Find()

	return c.Status(fiber.StatusFound).JSON(facts)
}

func (f *factHandler) CreateFact(c *fiber.Ctx) error {
	fact := dto.FactInput{}

	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	factOutput := f.factRepository.Create(&fact)

	return c.Status(fiber.StatusCreated).JSON(factOutput)
}
