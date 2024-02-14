package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gaspartv/GO-thygas-coins-back/internal/database"
	"github.com/gaspartv/GO-thygas-coins-back/internal/middlewares"
	"github.com/gaspartv/GO-thygas-coins-back/internal/service"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

type Env struct {
	DatabaseURL  string `validate:"required"`
	Port         string `validate:"required"`
	JwtSecretKey string `validate:"required"`
}

func validateEnv(env Env) error {
	validate := validator.New()
	if err := validate.Struct(env); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	var env Env
	env.DatabaseURL = os.Getenv("DATABASE_URL")
	env.Port = os.Getenv("PORT")
	env.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")

	if err := validateEnv(env); err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
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

	authService := service.NewAuthService(db)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.With(middlewares.JwtMiddleware).Post("/account-loyalty", accLoyaltyService.Create)
	c.Get("/account-loyalty/{id}", accLoyaltyService.Get)
	c.With(middlewares.JwtMiddleware).Delete("/account-loyalty/{id}", accLoyaltyService.Delete)
	c.Get("/account-loyalty", accLoyaltyService.List)
	c.With(middlewares.JwtMiddleware).Put("/account-loyalty/{id}", accLoyaltyService.Update)

	c.With(middlewares.JwtMiddleware).Post("/character", characterService.Create)
	c.Get("/character/{id}", characterService.Get)
	c.With(middlewares.JwtMiddleware).Delete("/character/{id}", characterService.Delete)
	c.Get("/character", characterService.List)
	c.With(middlewares.JwtMiddleware).Put("/character/{id}", characterService.Update)

	c.With(middlewares.JwtMiddleware).Post("/store", storeService.Create)
	c.Get("/store/{id}", storeService.Get)
	c.With(middlewares.JwtMiddleware).Delete("/store/{id}", storeService.Delete)
	c.Get("/store", storeService.List)
	c.With(middlewares.JwtMiddleware).Put("/store/{id}", storeService.Update)

	c.With(middlewares.JwtMiddleware).Post("/tibia-coins", tibiaCoinsService.Create)
	c.Get("/tibia-coins/{id}", tibiaCoinsService.Get)
	c.With(middlewares.JwtMiddleware).Delete("/tibia-coins/{id}", tibiaCoinsService.Delete)
	c.Get("/tibia-coins", tibiaCoinsService.List)
	c.With(middlewares.JwtMiddleware).Put("/tibia-coins/{id}", tibiaCoinsService.Update)

	c.With(middlewares.JwtMiddleware).Post("/promotion", promotionService.Create)
	c.Get("/promotion/{id}", promotionService.Get)
	c.With(middlewares.JwtMiddleware).Delete("/promotion/{id}", promotionService.Delete)
	c.Get("/promotion", promotionService.List)
	c.With(middlewares.JwtMiddleware).Put("/promotion/{id}", promotionService.Update)

	c.Post("/login", authService.Login)

	fmt.Println("Server is running on port", os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), c); err != nil {
		panic(err)
	}
}
