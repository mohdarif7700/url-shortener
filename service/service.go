package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/url-shortener/models"
)

func CreateShortenURL(ctx *fiber.Ctx) (models.ShortenURLResponse, error) {
	return models.ShortenURLResponse{}, nil
}
