package project

import (
	"demo-base/internal/service"

	"github.com/gofiber/fiber/v2"
)

func CreateRole(c *fiber.Ctx) error {
	role := service.RoleProject{}
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": nil,
			"msg":  "Invalid request body",
			"code": 400,
		})
	}
	if err := role.Create(c.Params("project")); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": nil,
			"msg":  err.Error(),
			"code": 500,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": role,
		"msg":  "success",
		"code": 200,
	})
}

func UpdateRole(c *fiber.Ctx) error {
	if c.Params("roleproject") == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": nil,
			"msg":  "Invalid request body",
			"code": 400,
		})
	}
	role := service.RoleProject{}
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": nil,
			"msg":  "Invalid request body",
			"code": 400,
		})
	}
	if err := role.Update(c.Params("project")); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": nil,
			"msg":  err.Error(),
			"code": 500,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": role,
		"msg":  "success",
		"code": 200,
	})
}

func DeleteRole(c *fiber.Ctx) error {
	role := service.RoleProject{
		Name: c.Params("roleproject"),
	}
	if err := role.Delete(c.Params("project")); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data": nil,
			"msg":  err.Error(),
			"code": 500,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": nil,
		"msg":  "success",
		"code": 200,
	})
}
