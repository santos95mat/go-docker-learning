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

func (*factRepository) Create(input *dto.FactInput) (*dto.FactOutput, error) {
	fact := model.Fact{
		Question: input.Question,
		Answer:   input.Answer,
	}

	err := database.DB.Create(&fact).Error

	if err != nil {
		return nil, err
	}

	return &dto.FactOutput{
		ID:        fact.ID,
		Question:  fact.Question,
		Answer:    fact.Answer,
		CreatedAt: fact.CreatedAt,
		UpdatedAt: fact.UpdatedAt,
	}, nil
}

func (*factRepository) FindAll() ([]*dto.FactOutput, error) {
	facts := []model.Fact{}
	factsOutput := []*dto.FactOutput{}

	err := database.DB.Find(&facts).Error

	if err != nil {
		return nil, err
	}

	for _, fact := range facts {
		factsOutput = append(factsOutput, &dto.FactOutput{
			ID:        fact.ID,
			Question:  fact.Question,
			Answer:    fact.Answer,
			CreatedAt: fact.CreatedAt,
			UpdatedAt: fact.UpdatedAt,
		})
	}

	return factsOutput, nil
}

func (*factRepository) FindOne(id int) (*dto.FactOutput, error) {
	fact := model.Fact{}

	err := database.DB.First(&fact, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &dto.FactOutput{
		ID:        fact.ID,
		Question:  fact.Question,
		Answer:    fact.Answer,
		CreatedAt: fact.CreatedAt,
		UpdatedAt: fact.UpdatedAt,
	}, nil
}
