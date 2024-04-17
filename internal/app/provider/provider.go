package provider

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-dashboard-api/internal/app/config"
	"github.com/jmoiron/sqlx"
)

type Provider struct {
	App       *fiber.App
	Config    *config.Config
	DB        *sqlx.DB
	Validator *config.Validator
}

func (p *Provider) Provide() {
}
