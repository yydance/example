package handler

import (
	"demo-base/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	login := &service.LoginInput{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	token, err := login.Login()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": 500,
			"data": nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "success",
		"code": 200,
		"data": token,
	})
}
