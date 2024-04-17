package provider

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-dashboard-api/internal/app/config"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/akun"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/auth"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/file"
	"github.com/jmoiron/sqlx"
)

type Provider struct {
	App       *fiber.App
	Config    *config.Config
	DB        *sqlx.DB
	Validator *config.Validator
}

func (p *Provider) Provide() {
	akun.ProvideAkun(p.App, p.Config, p.DB, p.Validator)
	auth.ProvideAuth(p.App, p.Config, p.DB, p.Validator)
	file.ProvideFile(p.App, p.Config, p.Validator)
}
