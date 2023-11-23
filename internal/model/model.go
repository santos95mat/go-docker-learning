package model

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null"`
	Answer   string `json:"answer" gorm:"text;not null;default:null"`
}

func NewFact(q string, a string) *Fact {
	return &Fact{
		Question: q,
		Answer:   a,
	}
}
