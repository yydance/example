package permission

import (
	"demo-base/internal/service"

	"github.com/gofiber/fiber/v2"
)

func ListRolePermission(c *fiber.Ctx) error {

	return c.Status(200).JSON(map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": service.ListRolePermission(),
	})
}

func ListProjectPermission(c *fiber.Ctx) error {

	return c.Status(200).JSON(map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": service.ListProjectPermission(),
	})
}
