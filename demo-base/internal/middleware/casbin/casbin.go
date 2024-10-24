package casbin

import (
	"demo-base/internal/models"

	"github.com/gofiber/fiber/v2"
)

func RoutePermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.FormValue("username")
		if user == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"msg":  "Unauthorized",
				"code": 401,
				"data": nil,
			})
		}
		res, err := models.CasbinEnforcer.Enforce(user, c.Path(), c.Method())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg":  "CasbinEnforcer.Enforce failed",
				"code": 500,
				"data": nil,
			})
		}
		if !res {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"msg":  "Forbidden",
				"code": 403,
				"data": nil,
			})
		}
		return c.Next()
	}
}
