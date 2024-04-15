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
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err.Error()})
	}

	resp, err := service.CreateShortenURL(req)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "data": err})
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "message": resp})
}

func RedirectURLHandler(ctx *fiber.Ctx) error {
	req := models.RedirectURLRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err.Error()})
	}

	resp, err := service.RedirectURL(req)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "data": err})
	}

	return ctx.Status(301).JSON(fiber.Map{"status": "success", "redirectingToURL": resp})
}

func GetMetricsHandler(ctx *fiber.Ctx) error {
	resp, err := service.GetMetrics()
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "data": "not found"})
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "success", "metrics": resp})
}
