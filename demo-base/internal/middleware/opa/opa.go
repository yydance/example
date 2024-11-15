package opa

import (
	"context"
	"demo-base/internal/utils/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage/inmem"
)

func OPA() fiber.Handler {

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
		_ = genPolicy(claims.UserName)
		ctx := context.Background()
		input_data := map[string]interface{}{
			"username": claims.UserName,
			"action":   c.Method(),
			"domain":   "global",
			"object":   c.Path(),
		}
		//store := inmem.NewFromReader(bytes.NewBufferString(policies))
		store := inmem.NewFromObject(policies)
		r := rego.New(
			rego.Query("data.authz.allow"),
			rego.Module("example.rego", module),
			rego.Store(store),
			rego.Input(input_data),
		)
		res, err := r.Eval(ctx)
		if err != nil || len(res) != 1 || len(res[0].Expressions) != 1 {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg":  "Server Error",
				"code": 500,
				"data": nil,
			})
		}
		if res[0].Expressions[0].Value == true {
			return c.Next()
		}
		if c.Params("project") != "" {
			input_data := map[string]interface{}{
				"username": claims.UserName,
				"action":   c.Method(),
				"domain":   c.Params("project"),
				"object":   c.Path(),
			}
			store := inmem.NewFromObject(policies)
			r := rego.New(
				rego.Query("data.authz.allow"),
				rego.Module("example.rego", module),
				rego.Store(store),
				rego.Input(input_data),
			)
			res, err := r.Eval(ctx)
			if err != nil || len(res) != 1 || len(res[0].Expressions) != 1 {
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"msg":  "Server Error",
					"code": 500,
					"data": nil,
				})
			}
			if res[0].Expressions[0].Value == true {
				return c.Next()
			}
			if c.Method() == "GET" && isPlatformView(claims.UserName) {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg":  "Authorization failed",
			"code": 401,
			"data": nil,
		})
	}
}
