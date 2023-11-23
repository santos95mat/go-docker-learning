package dto

import "time"

type FactInput struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type FactOutput struct {
	ID        uint      `json:"id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
