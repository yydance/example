package user

import (
	"demo-base/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	user := service.UserInput{}
	if err := c.BodyParser(user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":  "Invalid request body",
			"code": fiber.StatusBadRequest, // should be customizable code
			"data": nil,
		})
	}

	if err := user.Create(); err != nil {
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

func Get(c *fiber.Ctx) error {
	user := service.UserInput{
		Name: c.Params("username"),
	}

	u, err := user.Get()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": fiber.StatusNotFound,
			"data": nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Success",
		"code": fiber.StatusOK,
		"data": u,
	})
}

func Update(c *fiber.Ctx) error {
	user := service.UserInput{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": fiber.StatusBadRequest,
			"data": nil,
		})
	}
	if err := user.Update(); err != nil {
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

func UpdatePassword(c *fiber.Ctx) error {
	user := service.UserPassword{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": fiber.StatusBadRequest,
			"data": nil,
		})
	}
	if err := user.UpdatePassword(); err != nil {
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

func List(c *fiber.Ctx) error {
	user := service.UserInput{}
	users, err := user.List(c.QueryInt("page_num", 1), c.QueryInt("page_size", 10))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": fiber.StatusInternalServerError,
			"data": nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Success",
		"data": users,
		"code": fiber.StatusOK,
	})
}

func Delete(c *fiber.Ctx) error {
	user := service.UserInput{
		Name: c.Params("username"),
	}

	if err := user.Delete(); err != nil {
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
