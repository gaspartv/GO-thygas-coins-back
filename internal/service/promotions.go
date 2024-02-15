package service

import (
	"encoding/json"
	"net/http"

	"github.com/gaspartv/GO-thygas-coins-back/internal/database"
	"github.com/gaspartv/GO-thygas-coins-back/internal/entity"
	"github.com/gaspartv/GO-thygas-coins-back/internal/handlerError"
	"github.com/go-chi/chi/v5"
)

type PromotionService struct {
	db database.PromotionDB
}

func NewPromotionService(db database.PromotionDB) *PromotionService {
	return &PromotionService{
		db: db,
	}
}

func (service *PromotionService) Create(w http.ResponseWriter, r *http.Request) {
	var promotion entity.Promotion

	if err := json.NewDecoder(r.Body).Decode(&promotion); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if promotion.Description == "" {
		handlerError.Exec(w, "param description is required", http.StatusBadRequest)
		return
	}
	if promotion.Min == 0 {
		handlerError.Exec(w, "param min is required", http.StatusBadRequest)
		return
	}
	if promotion.Max == 0 {
		handlerError.Exec(w, "param max is required", http.StatusBadRequest)
		return
	}
	if promotion.Stack == 0 {
		handlerError.Exec(w, "param stack is required", http.StatusBadRequest)
		return
	}
	if promotion.Price == 0 {
		handlerError.Exec(w, "param price is required", http.StatusBadRequest)
		return
	}

	p := entity.NewPromotion(promotion.Description, promotion.Min, promotion.Max, promotion.Price, promotion.Stack)
	result, err := service.db.Create(p)
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

func (service *PromotionService) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	result, err := service.db.Get(id)
	if err != nil {
		handlerError.Exec(w, "promotion not found", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (service *PromotionService) Update(w http.ResponseWriter, r *http.Request) {
	var promotion entity.Promotion

	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&promotion); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if promotion.Description == "" {
		handlerError.Exec(w, "param description is required", http.StatusBadRequest)
		return
	}

	if promotion.Min == 0 {
		handlerError.Exec(w, "param min is required", http.StatusBadRequest)
		return
	}

	if promotion.Max == 0 {
		handlerError.Exec(w, "param max is required", http.StatusBadRequest)
		return
	}

	if promotion.Price == 0 {
		handlerError.Exec(w, "param price is required", http.StatusBadRequest)
		return
	}

	result, err := service.db.Update(id, &promotion)
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

func (service *PromotionService) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	if _, err := service.db.Get(id); err != nil {
		handlerError.Exec(w, "promotion not found", http.StatusNotFound)
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

func (service *PromotionService) List(w http.ResponseWriter, r *http.Request) {
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
