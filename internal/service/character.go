package service

import (
	"encoding/json"
	"net/http"

	"github.com/gaspartv/GO-thygas-coins-back/internal/database"
	"github.com/gaspartv/GO-thygas-coins-back/internal/entity"
	"github.com/gaspartv/GO-thygas-coins-back/internal/handlerError"
	"github.com/go-chi/chi/v5"
)

type CharacterService struct {
	db database.CharacterDB
}

func NewCharacterService(db database.CharacterDB) *CharacterService {
	return &CharacterService{
		db: db,
	}
}

func (handler *CharacterService) Create(w http.ResponseWriter, r *http.Request) {
	var character entity.Character

	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if character.Vocation == "" {
		handlerError.Exec(w, "param vocation is required", http.StatusBadRequest)
		return
	}
	if character.Level == 0 {
		handlerError.Exec(w, "param level is required", http.StatusBadRequest)
		return
	}

	if character.World == "" {
		handlerError.Exec(w, "param world is required", http.StatusBadRequest)
		return
	}

	if character.Description == "" {
		handlerError.Exec(w, "param description is required", http.StatusBadRequest)
		return
	}

	char := entity.NewCharacter(character.Vocation, character.Level, character.World, character.Description)
	result, err := handler.db.Create(char)
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

func (handler *CharacterService) Update(w http.ResponseWriter, r *http.Request) {
	var character entity.Character

	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if character.Vocation == "" {
		handlerError.Exec(w, "param vocation is required", http.StatusBadRequest)
		return
	}

	if character.Level == 0 {
		handlerError.Exec(w, "param level is required", http.StatusBadRequest)
		return
	}

	if character.World == "" {
		handlerError.Exec(w, "param world is required", http.StatusBadRequest)
		return
	}

	if character.Description == "" {
		handlerError.Exec(w, "param description is required", http.StatusBadRequest)
		return
	}

	result, err := handler.db.Update(id, &character)
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

func (handler *CharacterService) List(w http.ResponseWriter, r *http.Request) {
	result, err := handler.db.List()
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

func (handler *CharacterService) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	result, err := handler.db.Get(id)
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *CharacterService) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		handlerError.Exec(w, "param id is required", http.StatusBadRequest)
		return
	}

	if _, err := handler.db.Get(id); err != nil {
		handlerError.Exec(w, "character not found", http.StatusNotFound)
		return
	}

	result, err := handler.db.Delete(id)
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
