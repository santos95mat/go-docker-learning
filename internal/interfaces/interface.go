package interfaces

import (
	"github.com/santos95mat/go-docker-learning/internal/dto"
)

type FactRepoInterface interface {
	Find() ([]*dto.FactOutput, error)
	Create(fact *dto.FactInput) (*dto.FactOutput, error)
}
