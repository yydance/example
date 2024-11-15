package project

import (
	"demo-base/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	project := service.ProjectInput{}
	if err := c.BodyParser(&project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":   nil,
			"msg":    "Invalid request body",
			"status": fiber.StatusBadRequest,
		})
	}
	if err := project.Create(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":   nil,
			"msg":    err.Error(),
			"status": fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   nil,
		"msg":    "Project created successfully",
		"status": fiber.StatusOK,
	})
}

func Update(c *fiber.Ctx) error {
	project := service.ProjectInput{}
	if err := c.BodyParser(&project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":   nil,
			"msg":    "Invalid request body",
			"status": fiber.StatusBadRequest,
		})
	}
	if err := project.Update(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":   nil,
			"msg":    err.Error(),
			"status": fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   nil,
		"msg":    "Project updated successfully",
		"status": fiber.StatusOK,
	})
}

func Delete(c *fiber.Ctx) error {
	project := service.ProjectInput{
		Name: c.Params("project"),
	}
	if err := project.Delete(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":   nil,
			"msg":    err.Error(),
			"status": fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   nil,
		"msg":    "Project deleted successfully",
		"status": fiber.StatusOK,
	})
}

func List(c *fiber.Ctx) error {
	project := service.ProjectInput{}
	projects, err := project.List(c.QueryInt("page_num", 1), c.QueryInt("page_size", 10))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"data":   nil,
			"msg":    err.Error(),
			"status": fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   projects,
		"msg":    "Project list fetched successfully",
		"status": fiber.StatusOK,
	})
}
