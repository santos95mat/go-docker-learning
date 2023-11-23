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
	facts := f.factRepository.Find()

	return c.Status(fiber.StatusFound).JSON(facts)
}

func (f *factHandler) CreateFact(c *fiber.Ctx) error {
	factInput := dto.FactInput{}

	if err := c.BodyParser(&factInput); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fact := f.factRepository.Create(&factInput)

	return c.Status(fiber.StatusCreated).JSON(fact)
}
