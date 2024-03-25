package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/tnnz20/final-hacktiv8-fga/config"
	handler "github.com/tnnz20/final-hacktiv8-fga/internal/adapters/handler/http"
	"github.com/tnnz20/final-hacktiv8-fga/internal/adapters/handler/http/router"
	"github.com/tnnz20/final-hacktiv8-fga/internal/adapters/storage/postgres"
	"github.com/tnnz20/final-hacktiv8-fga/internal/adapters/storage/postgres/repository"
	"github.com/tnnz20/final-hacktiv8-fga/internal/core/service"
)

func main() {
	// Load yaml config
	cfg, err := config.LoadConfig("local")
	if err != nil {
		panic(err)
	}

	var (
		CfgPostgres  = cfg.GetDatabaseConfig().Postgres
		JwtSecretKey = cfg.GetTokenConfig().JWTSecret
	)

	// Connection database
	db, err := postgres.New(&CfgPostgres)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err := db.Migrate(); err != nil {
		panic(err)
	}

	// Create new echo instance
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "- method=${method}, uri=${uri}, status=${status}\n",
	}))
	// e.Use(middleware.Recover())

	// Validator
	validate := validator.New()

	// User
	userRepo := repository.NewUserRepository(db.GetDB())
	userService := service.NewUserService(userRepo, &JwtSecretKey)
	userHandler := handler.NewUserHandler(userService, validate)

	router.NewUserRouter(e, &JwtSecretKey, userHandler)

	// Start server
	cfgServer := cfg.GetServerConfig()
	PORT := fmt.Sprint(cfgServer.Port)
	if PORT == "" {
		PORT = "8080"
	}
	s := http.Server{
		Addr:    ":" + PORT,
		Handler: e,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}

	fmt.Println("Server is running on port", PORT)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
