package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/url-shortener/handler"
)

func SetupRoutes(app *fiber.App) {
	// API endpoints
	app.Post("/shorten-url", handler.ShortenURLHandler)
	//app.Get("/:shortURL", RedirectURL)
	//app.Get("/metrics", GetMetrics)
}
