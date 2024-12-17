package controllerr

import (
	"github.com/gofiber/fiber/v2"
)

func IndexController(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"error_code": "0",
		"message":    "Welcome to shortner",
	})
}
