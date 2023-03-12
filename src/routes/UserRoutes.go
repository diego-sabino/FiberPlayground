package routes

import (
	"FiberPlayground/src/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/users", controller.GetAll)
	app.Get("/users/:id", controller.GetById)
	app.Post("/users", controller.Create)
	app.Put("/users/:id", controller.Update)
	app.Delete("/users/:id", controller.Delete)

	return app
}
