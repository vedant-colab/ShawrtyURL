package routes

import (
	controller "src/controllers"

	"github.com/gofiber/fiber/v2"
)

func IndexRouter(router fiber.Router) {
	router.Get("/", controller.IndexController)
}
