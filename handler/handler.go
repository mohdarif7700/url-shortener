package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/url-shortener/models"
	"github.com/url-shortener/service"
)

func ShortenURLHandler(ctx *fiber.Ctx) error {
	req := models.ShortenURLRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	resp, err := service.CreateShortenURL(ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "data": err})
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": resp})
}
