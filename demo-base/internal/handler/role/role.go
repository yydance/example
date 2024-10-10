package role

import (
	"demo-base/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	role := service.RoleInput{}
	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":  "Invalid request body",
			"code": fiber.StatusBadRequest,
			"data": nil,
		})
	}

	if err := role.Create(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": fiber.StatusInternalServerError,
			"data": nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "success",
		"code": fiber.StatusOK,
		"data": nil,
	})
}

func Update(c *fiber.Ctx) error {
	role := service.RoleInput{}
	if err := c.BodyParser(role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":  "Invalid request body",
			"code": fiber.StatusBadRequest,
			"data": nil,
		})
	}

	if err := role.Update(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": fiber.StatusInternalServerError,
			"data": nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Success",
		"code": fiber.StatusOK,
		"data": nil,
	})
}
