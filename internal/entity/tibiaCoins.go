package entity

import "github.com/google/uuid"

type TibiaCoins struct {
	ID     string  `json:"id"`
	Code   string  `json:"code"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount int     `json:"amount"`
	Min    int     `json:"min"`
	Max    int     `json:"max"`
	Image  string  `json:"image"`
	Step   int     `json:"step"`
}

func NewTibiaCoins(code string, name string, price float64, amount int, min int, max int, image string, step int) *TibiaCoins {
	return &TibiaCoins{
		ID:     uuid.New().String(),
		Code:   code,
		Name:   name,
		Price:  price,
		Amount: amount,
		Min:    min,
		Max:    max,
		Image:  image,
		Step:   step,
	}
}
