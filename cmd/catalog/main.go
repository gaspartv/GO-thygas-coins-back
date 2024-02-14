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

	characterDB := database.NewCharacterDB(db)
	characterService := service.NewCharacterService(*characterDB)

	storeDB := database.NewStoreDB(db)
	storeService := service.NewStoreService(*storeDB)

	tibiaCoinsDB := database.NewTibiaCoinsDB(db)
	tibiaCoinsService := service.NewTibiaCoinsService(*tibiaCoinsDB)

	promotionDB := database.NewPromotionDB(db)
	promotionService := service.NewPromotionService(*promotionDB)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Post("/account-loyalty", accLoyaltyService.Create)
	c.Get("/account-loyalty/{id}", accLoyaltyService.Get)
	c.Delete("/account-loyalty/{id}", accLoyaltyService.Delete)
	c.Get("/account-loyalty", accLoyaltyService.List)
	c.Put("/account-loyalty/{id}", accLoyaltyService.Update)

	c.Post("/character", characterService.Create)
	c.Get("/character/{id}", characterService.Get)
	c.Delete("/character/{id}", characterService.Delete)
	c.Get("/character", characterService.List)
	c.Put("/character/{id}", characterService.Update)

	c.Post("/store", storeService.Create)
	c.Get("/store/{id}", storeService.Get)
	c.Delete("/store/{id}", storeService.Delete)
	c.Get("/store", storeService.List)
	c.Put("/store/{id}", storeService.Update)

	c.Post("/tibia-coins", tibiaCoinsService.Create)
	c.Get("/tibia-coins/{id}", tibiaCoinsService.Get)
	c.Delete("/tibia-coins/{id}", tibiaCoinsService.Delete)
	c.Get("/tibia-coins", tibiaCoinsService.List)
	c.Put("/tibia-coins/{id}", tibiaCoinsService.Update)

	c.Post("/promotion", promotionService.Create)
	c.Get("/promotion/{id}", promotionService.Get)
	c.Delete("/promotion/{id}", promotionService.Delete)
	c.Get("/promotion", promotionService.List)
	c.Put("/promotion/{id}", promotionService.Update)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", c); err != nil {
		panic(err)
	}
}
