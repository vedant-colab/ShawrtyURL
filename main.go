package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"src/routes"

	// "src\routes"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer dbpool.Close()
	if err := dbpool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	fmt.Println("Connected to PostgreSQL database!")

	baseUrl := []string{os.Getenv("BaseURL"), os.Getenv("version")}
	urls := strings.Join(baseUrl, "/")

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/" + urls)
	routes.IndexRouter(api)
	app.Listen(":8080")
}
