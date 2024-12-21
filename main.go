package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ThisIsNotJustin/goshort/api/database"
	"github.com/ThisIsNotJustin/goshort/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading env file: ", err)
	}

	database.InitializeClient()

	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

func setupRoutes(app *fiber.App) {
	app.Post("/api/v1", routes.ShortenURL)
	app.Get("/api/v1/:shortID", routes.GetByShortID)
	app.Put("/api/v1/:shortID", routes.EditURL)
	app.Delete("api/v1/:shortID", routes.DeleteURL)
	app.Get("/:shortID", routes.GetByShortID)
}
