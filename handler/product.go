package handler

import (
	"fiber-101/model"

	"github.com/gofiber/fiber/v2"
)

func LineWebhooks(c *fiber.Ctx) error {
	var product []model.Product

	return c.JSON(fiber.Map{
		"data": product,
	})
}
