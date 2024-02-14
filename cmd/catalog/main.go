package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gaspartv/GO-thygas-coins-back/internal/database"
	"github.com/gaspartv/GO-thygas-coins-back/internal/service"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/thygas_coins")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	accountLoyaltyDB := database.NewAccLoyaltyDB(db)
	accLoyaltyService := service.NewAccLoyaltyService(*accountLoyaltyDB)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Post("/account-loyalty", accLoyaltyService.Create)
	c.Get("/account-loyalty/{id}", accLoyaltyService.Get)
	c.Delete("/account-loyalty/{id}", accLoyaltyService.Delete)
	c.Get("/account-loyalty", accLoyaltyService.List)
	c.Put("/account-loyalty/{id}", accLoyaltyService.Update)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", c); err != nil {
		panic(err)
	}
}
