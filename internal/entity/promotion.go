package entity

import "github.com/google/uuid"

type Promotion struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Min         int     `json:"min"`
	Max         int     `json:"max"`
	Price       float64 `json:"price"`
}

func NewPromotion(description string, min int, max int, price float64) *Promotion {
	return &Promotion{
		ID:          uuid.New().String(),
		Description: description,
		Min:         min,
		Max:         max,
		Price:       price,
	}
}
