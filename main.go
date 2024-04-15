package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/url-shortener/router"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	router.SetupRoutes(app)

	// Start the server
	app.Listen(":8080")

}
