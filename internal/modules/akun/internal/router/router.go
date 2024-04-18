package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikti-its/khanza-dashboard-api/internal/app/middleware"
	"github.com/ikti-its/khanza-dashboard-api/internal/modules/akun/internal/controller"
)

func Route(
	app *fiber.App,
	controller *controller.AkunController,
) {
	akun := app.Group("/v1/akun")
	{
		akun.Post("/", middleware.Authenticate([]int{1000}), controller.Create)
		akun.Get("/", middleware.Authenticate([]int{1000, 2000}), controller.Get)
		akun.Get("/:id", middleware.Authenticate([]int{1000, 2000}), controller.GetById)
		akun.Put("/:id", middleware.Authenticate([]int{1000, 2000}), controller.Update)
		akun.Delete("/:id", middleware.Authenticate([]int{1000}), controller.Delete)
	}
}
