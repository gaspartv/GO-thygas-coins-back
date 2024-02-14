package service

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gaspartv/GO-thygas-coins-back/internal/entity"
	"github.com/gaspartv/GO-thygas-coins-back/internal/handlerError"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db *sql.DB
}

type TokenMessage struct {
	Token string `json:"token"`
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (handler *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	var auth entity.Auth

	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	if auth.Email == "" {
		handlerError.Exec(w, "param email is required", http.StatusBadRequest)
		return
	}

	if auth.Password == "" {
		handlerError.Exec(w, "param password is required", http.StatusBadRequest)
		return
	}

	stmt, err := handler.db.Prepare("SELECT * FROM stores WHERE email =?")
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(auth.Email)

	var store entity.Store
	if err := row.Scan(&store.ID, &store.Name, &store.QRCode, &store.Email, &store.Cellphone, &store.Password); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	decodedHash, err := base64.StdEncoding.DecodeString(store.Password)
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword(decodedHash, []byte(auth.Password))
	if err != nil {
		handlerError.Exec(w, "email or password incorrect", http.StatusBadRequest)
		return
	}

	claims := jwt.MapClaims{
		"username": store.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := TokenMessage{
		Token: tokenString,
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		handlerError.Exec(w, err.Error(), http.StatusBadRequest)
		return
	}
}
