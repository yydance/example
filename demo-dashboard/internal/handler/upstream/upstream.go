package upstream

import (
	"demo-dashboard/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	appHandler := handler.Fiber{C: c}

	return nil
}
