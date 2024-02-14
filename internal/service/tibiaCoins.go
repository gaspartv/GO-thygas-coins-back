package service

import (
	"encoding/json"
	"net/http"

	"github.com/gaspartv/GO-thygas-coins-back/internal/database"
	"github.com/gaspartv/GO-thygas-coins-back/internal/entity"
	"github.com/gaspartv/GO-thygas-coins-back/internal/handlerError"
	"github.com/go-chi/chi/v5"
)

type TibiaCoinsService struct {
	db database.TibiaCoinsDB
}

func NewTibiaCoinsService(db database.TibiaCoinsDB) *TibiaCoinsService {
	return &TibiaCoinsService{
		db: db,
	}
}

func (service *TibiaCoinsService) Create(w http.ResponseWriter, r *http.Request) {
	var tibiaCoins entity.TibiaCoins

	if err := json.NewDecoder(r.Body).Decode(&tibiaCoins); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if tibiaCoins.Code == "" {
		handlerError.Exec(w, "code is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Name == "" {
		handlerError.Exec(w, "name is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Price == 0 {
		handlerError.Exec(w, "price is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Amount == 0 {
		handlerError.Exec(w, "amount is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Min == 0 {
		handlerError.Exec(w, "min is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Max == 0 {
		handlerError.Exec(w, "max is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Image == "" {
		handlerError.Exec(w, "image is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Step == 0 {
		handlerError.Exec(w, "step is required", http.StatusBadRequest)
		return
	}

	tc := entity.NewTibiaCoins(tibiaCoins.Code, tibiaCoins.Name, tibiaCoins.Price, tibiaCoins.Amount, tibiaCoins.Min, tibiaCoins.Max, tibiaCoins.Image, tibiaCoins.Step)
	result, err := service.db.Create(tc)
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := handlerError.Response{
		Message: result,
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (service *TibiaCoinsService) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	result, err := service.db.Get(id)
	if err != nil {
		handlerError.Exec(w, "tibia coins not found", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (service *TibiaCoinsService) List(w http.ResponseWriter, r *http.Request) {
	result, err := service.db.List()
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (service *TibiaCoinsService) Update(w http.ResponseWriter, r *http.Request) {
	var tibiaCoins entity.TibiaCoins

	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&tibiaCoins); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if tibiaCoins.Code == "" {
		handlerError.Exec(w, "code is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Name == "" {
		handlerError.Exec(w, "name is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Price == 0 {
		handlerError.Exec(w, "price is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Amount == 0 {
		handlerError.Exec(w, "amount is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Min == 0 {
		handlerError.Exec(w, "min is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Max == 0 {
		handlerError.Exec(w, "max is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Image == "" {
		handlerError.Exec(w, "image is required", http.StatusBadRequest)
		return
	}

	if tibiaCoins.Step == 0 {
		handlerError.Exec(w, "step is required", http.StatusBadRequest)
		return
	}

	if _, err := service.db.Get(id); err != nil {
		handlerError.Exec(w, "tibia coins not found", http.StatusNotFound)
		return
	}

	result, err := service.db.Update(id, &tibiaCoins)
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := handlerError.Response{
		Message: result,
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (service *TibiaCoinsService) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	if _, err := service.db.Get(id); err != nil {
		handlerError.Exec(w, "tibia coins not found", http.StatusNotFound)
		return
	}

	result, err := service.db.Delete(id)
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := handlerError.Response{
		Message: result,
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
