package entity

import "github.com/google/uuid"

type Character struct {
	ID          string `json:"id"`
	Vocation    string `json:"vocation"`
	Level       int    `json:"level"`
	World       string `json:"world"`
	Description string `json:"description"`
}

func NewCharacter(vocation string, level int, world string, description string) *Character {
	return &Character{
		ID:          uuid.New().String(),
		Vocation:    vocation,
		Level:       level,
		World:       world,
		Description: description,
	}
}
