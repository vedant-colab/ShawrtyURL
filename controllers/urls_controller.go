package controller

import (
	"src/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ShortenURLBody struct {
	ShortenURL string `json:"shortenURL" validate:"required,min=5"`
	ActualURL  string `json:"actualURL" validate:"required,min=5"`
}

// Validator instance
var validate = validator.New()

func ShortenUrl(c *fiber.Ctx) error {
	// Parse request body into struct
	var shortenurlBody ShortenURLBody
	if err := c.BodyParser(&shortenurlBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error_code":    "1",
			"error_message": "Invalid request body",
		})
	}

	// Validate request body
	if err := validate.Struct(&shortenurlBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error_code":    "2",
			"error_message": "Validation failed",
			"details":       err.Error(),
		})
	}

	// Call the service to save the URL
	if err := services.SaveShortenURL(shortenurlBody.ShortenURL, shortenurlBody.ActualURL); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error_code":    "3",
			"error_message": "Failed to save URL",
			"details":       err.Error(),
		})
	}

	// Success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error_code":    "0",
		"error_message": "URL saved successfully",
	})
}

func GetAllURLs(c *fiber.Ctx) error {
	rows, err := services.FetchAllURL()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error_code":    "1",
				"error_message": "couldn't fetch rows",
				"details":       err.Error(),
			})
		// return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error_code": "0",
		"Data":       rows,
	})
	// return nil
}
