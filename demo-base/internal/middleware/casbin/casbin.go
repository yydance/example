package casbin

import (
	"demo-base/internal/models"
	"demo-base/internal/utils/jwt"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gofiber/fiber/v2"
)

func RoutePermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Path() == "/live" || c.Path() == "/ready" || c.Path() == "/api/v1/login" {
			return c.Next()
		}
		// 解析token获取用户名，并鉴权
		var jwt *jwt.JWT
		claims, err := jwt.ParseToken(c.Get("Authorization"))
		if err != nil || claims == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code": 401,
				"msg":  err.Error(),
				"data": nil,
			})
		}
		if claims.UserName == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"msg":  "Unauthorized",
				"code": 401,
				"data": nil,
			})
		}
		// 获取role，如果是管理员Admin，则直接放行
		roles, err := models.CasbinEnforcer.GetRolesForUser(claims.UserName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg":  "CasbinEnforcer.GetRolesForUser failed",
				"code": 500,
				"data": nil,
			})
		}
		if len(roles) > 0 {
			mapsets := mapset.NewSet[string]()
			for _, role := range roles {
				mapsets.Add(role)
			}
			if mapsets.Contains("Admin") {
				return c.Next()
			}
		}

		// 检查普通用户权限
		if strings.Contains(c.Path(), "/api/v1/projects") && c.Params("name") != "" {
			res, err := models.CasbinEnforcer.Enforce(claims.UserName, c.Params("name"), c.Path(), c.Method())
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"msg":  "CasbinEnforcer.Enforce failed",
					"code": 500,
					"data": nil,
				})
			}
			if !res {
				res, err := models.CasbinEnforcer.Enforce(claims.UserName, "global", c.Path(), c.Method())
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
			}
		}
		return c.Next()
	}
}
