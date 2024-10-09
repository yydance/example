package system

import (
	"demo-base/internal/conf"

	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "pong",
		"data": nil,
	})
}

func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "ok",
		"data": nil,
	})
}

func Info(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"version": conf.Version,
		"build":   "2022-01-01",
	})
}
