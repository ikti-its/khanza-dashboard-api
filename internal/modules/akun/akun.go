package akun

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-dashboard-api/internal/app/config"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/akun/internal/controller"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/akun/internal/repository/postgres"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/akun/internal/router"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/akun/internal/usecase"
	"github.com/jmoiron/sqlx"
)

func ProvideAkun(app *fiber.App, cfg *config.Config, db *sqlx.DB, validator *config.Validator) {
	akunRepository := postgres.NewAkunRepository(db)
	akunUseCase := usecase.NewAkunUseCase(&akunRepository, cfg)
	akunController := controller.NewAkunController(akunUseCase, validator)

	router.Route(
		app,
		akunController,
	)
}
