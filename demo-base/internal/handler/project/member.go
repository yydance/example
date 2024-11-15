package project

import (
	"demo-base/internal/service"

	"github.com/gofiber/fiber/v2"
)

func AddMember(c *fiber.Ctx) error {
	member := service.ProjectMemberInput{}
	if err := c.BodyParser(&member); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Invalid request body",
		})
	}
	if err := member.AddMember(c.Params("project")); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":  err.Error(),
			"data": nil,
			"code": 500,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "success",
		"data": nil,
		"code": 200,
	})
}

func RemoveMember(c *fiber.Ctx) error {
	member := service.ProjectMemberInput{
		UserName: c.Params("memberName"),
	}
	if err := member.DeleteMember(c.Params("project")); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":  err.Error(),
			"data": nil,
			"code": 500,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "success",
		"data": nil,
		"code": 200,
	})
}

func UpdateMember(c *fiber.Ctx) error {
	member := service.ProjectMemberInput{}
	if err := c.BodyParser(&member); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Invalid request body",
		})
	}
	if err := member.UpdateMember(c.Params("project")); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":  err.Error(),
			"data": nil,
			"code": 500,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "success",
		"data": nil,
		"code": 200,
	})
}
