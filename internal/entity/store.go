package entity

import "github.com/google/uuid"

type Store struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	QRCode    string `json:"qrcode"`
	Email     string `json:"email"`
	Cellphone string `json:"cellphone"`
	Password  string `json:"password"`
}

func NewStore(name string, qrCode string, email string, cellphone string, password string) *Store {
	return &Store{
		ID:        uuid.New().String(),
		Name:      name,
		QRCode:    qrCode,
		Email:     email,
		Cellphone: cellphone,
		Password:  password,
	}
}
