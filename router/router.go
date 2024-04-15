package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/url-shortener/handler"
)

func SetupRoutes(app *fiber.App) {
	// API endpoints
	app.Post("/shorten-url", handler.ShortenURLHandler)
	app.Post("/redirect", handler.RedirectURLHandler)
	app.Get("/metrics", handler.GetMetricsHandler)
}
