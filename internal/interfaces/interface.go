package interfaces

import (
	"github.com/santos95mat/go-docker-learning/internal/dto"
)

type FactRepoInterface interface {
	Create(fact *dto.FactInput) (*dto.FactOutput, error)
	FindAll() ([]*dto.FactOutput, error)
	FindOne(id int) (*dto.FactOutput, error)
}
