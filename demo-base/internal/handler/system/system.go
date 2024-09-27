package system

import (
	"demo-base/internal/conf"

	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}

func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}

func Info(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"version": conf.Version,
		"build":   "2022-01-01",
	})
}
