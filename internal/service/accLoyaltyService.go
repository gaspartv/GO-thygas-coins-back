package service

import (
	"encoding/json"
	"net/http"

	"github.com/gaspartv/GO-thygas-coins-back/internal/database"
	"github.com/gaspartv/GO-thygas-coins-back/internal/entity"
	"github.com/gaspartv/GO-thygas-coins-back/internal/handlerError"
	"github.com/go-chi/chi/v5"
)

type AccLoyaltyService struct {
	AccLoyaltyDB database.AccLoyaltyDB
}

func NewAccLoyaltyService(accountLoyaltyDB database.AccLoyaltyDB) *AccLoyaltyService {
	return &AccLoyaltyService{AccLoyaltyDB: accountLoyaltyDB}
}

func (handler *AccLoyaltyService) Create(w http.ResponseWriter, r *http.Request) {
	var accountLoyalty entity.AccLoyalty

	if err := json.NewDecoder(r.Body).Decode(&accountLoyalty); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if accountLoyalty.Percentage < 5 || accountLoyalty.Percentage > 50 {
		handlerError.Exec(w, "percentage should be between 5 to 50 ", http.StatusConflict)
		return
	}

	if accountLoyalty.Price <= 0 {
		handlerError.Exec(w, "price should be greater than 0 ", http.StatusConflict)
		return
	}

	accLoyalty := entity.NewAccLoyalty(accountLoyalty.Percentage, accountLoyalty.Price)
	result, err := handler.AccLoyaltyDB.Create(accLoyalty)
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

func (handler *AccLoyaltyService) Update(w http.ResponseWriter, r *http.Request) {
	var accountLoyalty entity.AccLoyalty

	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&accountLoyalty); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if accountLoyalty.Percentage < 5 || accountLoyalty.Percentage > 50 {
		handlerError.Exec(w, "percentage should be between 5 to 50 ", http.StatusConflict)
		return
	}

	if accountLoyalty.Price <= 0 {
		handlerError.Exec(w, "price should be greater than 0 ", http.StatusConflict)
		return
	}

	result, err := handler.AccLoyaltyDB.Update(id, &accountLoyalty)
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

func (handler *AccLoyaltyService) List(w http.ResponseWriter, r *http.Request) {
	result, err := handler.AccLoyaltyDB.List()
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

func (handler *AccLoyaltyService) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	result, err := handler.AccLoyaltyDB.Get(id)
	if err != nil {
		handlerError.Exec(w, "account loyalty not found", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *AccLoyaltyService) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	if _, err := handler.AccLoyaltyDB.Get(id); err != nil {
		handlerError.Exec(w, "account loyalty not found", http.StatusNotFound)
		return
	}

	result, err := handler.AccLoyaltyDB.Delete(id)
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := handlerError.Response{
		Message: result,
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}
}
