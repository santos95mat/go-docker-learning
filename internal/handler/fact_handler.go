package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-docker-learning/internal/dto"
	"github.com/santos95mat/go-docker-learning/internal/interfaces"
)

type factHandler struct {
	factRepository interfaces.FactRepoInterface
}

func NewFactHandler(repo interfaces.FactRepoInterface) *factHandler {
	return &factHandler{factRepository: repo}
}

func (f *factHandler) ListFacts(c *fiber.Ctx) error {
	facts, err := f.factRepository.Find()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusFound).JSON(facts)
}

func (f *factHandler) CreateFact(c *fiber.Ctx) error {
	factInput := dto.FactInput{}

	if err := c.BodyParser(&factInput); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fact, err := f.factRepository.Create(&factInput)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fact)
}
