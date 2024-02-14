package entity

import "github.com/google/uuid"

type AccLoyalty struct {
	ID         string  `json:"id"`
	Percentage int     `json:"percentage"`
	Price      float64 `json:"price"`
}

func NewAccLoyalty(percentage int, price float64) *AccLoyalty {
	return &AccLoyalty{
		ID:         uuid.New().String(),
		Percentage: percentage,
		Price:      price,
	}
}
