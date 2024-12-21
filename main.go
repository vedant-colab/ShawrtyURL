package main

import (
	"fmt"
	"log"
	"os"
	"src/database"
	"src/routes"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	defer database.DB.Close()

	fmt.Println("Connected to PostgreSQL database!")

	baseUrl := []string{os.Getenv("BaseURL"), os.Getenv("version")}
	urls := strings.Join(baseUrl, "/")

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/" + urls)
	routes.IndexRouter(api)
	routes.UrlRouter(api)
	app.Listen(":8080")
}
