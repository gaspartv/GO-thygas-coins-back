package entity

import "github.com/google/uuid"

type TibiaCoinsPromotion struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Min         int     `json:"min"`
	Max         int     `json:"max"`
	Price       float64 `json:"price"`
}

func NewTibiaCoinsPromotion(id string, description string, min int, max int, price float64) *TibiaCoinsPromotion {
	return &TibiaCoinsPromotion{
		ID:          uuid.New().String(),
		Description: description,
		Min:         min,
		Max:         max,
		Price:       price,
	}
}
