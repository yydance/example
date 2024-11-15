package jwt

import (
	"demo-base/internal/conf"
	"demo-base/internal/service"
	"demo-base/internal/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func Authentication() fiber.Handler {

	return func(c *fiber.Ctx) error {
		var jwt *jwt.JWT
		claims, err := jwt.ParseToken(c.Get("Authorization"))
		if err != nil || claims == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code": 401,
				"msg":  err.Error(),
				"data": nil,
			})
		}
		// 检验用户名，Issuer
		user := service.UserInput{Name: claims.UserName}
		if _, ok := user.IsExist(); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code": 401,
				"msg":  "用户不存在",
				"data": nil,
			})
		}
		if claims.Issuer != conf.Issuer {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code": 401,
				"msg":  "无效的token",
				"data": nil,
			})
		}
		return c.Next()
	}
}
