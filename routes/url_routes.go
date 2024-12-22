package routes

import (
	controller "src/controllers"

	"github.com/gofiber/fiber/v2"
)

func UrlRouter(router fiber.Router) {
	router.Post("/shorten-url", controller.ShortenUrl)
	router.Get("/get-urls", controller.GetAllURLs)
	router.Get("/get-url/:shortenURL", controller.GetURL)
}
