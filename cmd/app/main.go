package main

import (
	"fmt"
	"github.com/ikti-its/khanza-dashboard-api/internal/app/config"
	"github.com/ikti-its/khanza-dashboard-api/internal/app/provider"
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	var (
		cfg       = config.NewConfig()
		fiber     = config.NewFiber(cfg)
		postgres  = config.NewPostgres(cfg)
		validator = config.NewValidator()
		bootstrap = provider.Provider{App: fiber, Config: cfg, DB: postgres, Validator: validator}
	)

	defer func(postgres *sqlx.DB) {
		err := postgres.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}(postgres)

	bootstrap.Provide()

	if err := fiber.Listen(fmt.Sprintf(":%d", cfg.GetInt("APP_PORT", 8080))); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
