package interfaces

import (
	"github.com/santos95mat/go-docker-learning/internal/dto"
)

type FactRepository interface {
	Find() []*dto.FactOutput
	Create(fact *dto.FactInput) *dto.FactOutput
}
