package controller

import (
	"src/services"
	"src/utils"

	"github.com/gofiber/fiber/v2"
)

// Validator instance

func ShortenUrl(c *fiber.Ctx) error {
	// Parse request body into struct
	body := new(struct {
		ActualURL string `json:"actualURL" validate:"required"`
	})

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	shortenedURL, err := services.SaveShortenURL(body.ActualURL)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not save URL"})
	}

	finalUrl := utils.GenerateURL(shortenedURL)

	return c.Status(201).JSON(fiber.Map{
		"shorten_url": finalUrl,
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

func GetURL(c *fiber.Ctx) error {
	shortenURL := c.Params("shortenURL")
	if shortenURL == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	actualURL, err := services.FetchActualURL(shortenURL)
	if err != nil {
		// log.Fatalf("Error fetching all urls from database : %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch URL"})
	}
	return c.Status(200).JSON(fiber.Map{
		"actual_url": actualURL,
		"error_code": "0",
	})
}
