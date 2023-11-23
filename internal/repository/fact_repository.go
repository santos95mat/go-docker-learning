package repository

import (
	"github.com/santos95mat/go-docker-learning/internal/database"
	"github.com/santos95mat/go-docker-learning/internal/dto"
	"github.com/santos95mat/go-docker-learning/internal/model"
)

type factRepository struct {
}

func NewFactRepository() *factRepository {
	return &factRepository{}
}

func (*factRepository) Find() []*dto.FactOutput {
	facts := []model.Fact{}
	factsOutput := []*dto.FactOutput{}

	database.DB.Find(&facts)

	for _, fact := range facts {
		factsOutput = append(factsOutput, &dto.FactOutput{
			ID:        fact.ID,
			Question:  fact.Question,
			Answer:    fact.Answer,
			CreatedAt: fact.CreatedAt,
			UpdatedAt: fact.UpdatedAt,
		})
	}

	return factsOutput
}

func (*factRepository) Create(input *dto.FactInput) *dto.FactOutput {
	fact := model.Fact{
		Question: input.Question,
		Answer:   input.Answer,
	}

	database.DB.Create(&fact)

	return &dto.FactOutput{
		ID:        fact.ID,
		Question:  fact.Question,
		Answer:    fact.Answer,
		CreatedAt: fact.CreatedAt,
		UpdatedAt: fact.UpdatedAt,
	}
}
