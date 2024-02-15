package service

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/gaspartv/GO-thygas-coins-back/internal/database"
	"github.com/gaspartv/GO-thygas-coins-back/internal/entity"
	"github.com/gaspartv/GO-thygas-coins-back/internal/handlerError"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type StoreService struct {
	db database.StoreDB
}

func NewStoreService(db database.StoreDB) *StoreService {
	return &StoreService{
		db: db,
	}
}

func (service *StoreService) Create(w http.ResponseWriter, r *http.Request) {
	var store entity.Store

	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if store.Name == "" {
		handlerError.Exec(w, "name is required", http.StatusBadRequest)
		return
	}

	if store.QRCode == "" {
		handlerError.Exec(w, "qrcode is required", http.StatusBadRequest)
		return
	}

	if store.Email == "" {
		handlerError.Exec(w, "email is required", http.StatusBadRequest)
		return
	}

	if store.Cellphone == "" {
		handlerError.Exec(w, "cellphone is required", http.StatusBadRequest)
		return
	}

	if store.Password == "" {
		handlerError.Exec(w, "password is required", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(store.Password), bcrypt.DefaultCost)
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}

	passwordHash := base64.StdEncoding.EncodeToString(hash)

	s := entity.NewStore(store.Name, store.QRCode, store.Email, store.Cellphone, passwordHash)
	result, err := service.db.Create(s)
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

func (handler *StoreService) Update(w http.ResponseWriter, r *http.Request) {
	var store entity.Store

	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if store.Name == "" {
		handlerError.Exec(w, "name is required", http.StatusBadRequest)
		return
	}

	if store.QRCode == "" {
		handlerError.Exec(w, "qrcode is required", http.StatusBadRequest)
		return
	}

	if store.Email == "" {
		handlerError.Exec(w, "email is required", http.StatusBadRequest)
		return
	}

	if store.Cellphone == "" {
		handlerError.Exec(w, "cellphone is required", http.StatusBadRequest)
		return
	}

	if store.Password == "" {
		handlerError.Exec(w, "password is required", http.StatusBadRequest)
		return
	}

	if _, err := handler.db.Get(); err != nil {
		handlerError.Exec(w, "store not found", http.StatusNotFound)
		return
	}

	result, err := handler.db.Update(id, &store)
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

func (handler *StoreService) Get(w http.ResponseWriter, r *http.Request) {
	result, err := handler.db.Get()
	if err != nil {
		handlerError.Exec(w, "store not found", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *StoreService) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	if _, err := handler.db.Get(); err != nil {
		handlerError.Exec(w, "store not found", http.StatusNotFound)
		return
	}

	result, err := handler.db.Delete(id)
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
