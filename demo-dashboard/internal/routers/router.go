package routers

import (
	"demo-dashboard/internal/conf"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func InitRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:         conf.ServerOption.AppName,
		BodyLimit:       conf.ServerOption.BodyLimit,
		Concurrency:     conf.ServerOption.Concurrency,
		IdleTimeout:     conf.ServerOption.IdleTimeout,
		ReadTimeout:     conf.ServerOption.ReadTimeout,
		WriteTimeout:    conf.ServerOption.WriteTimeout,
		ReadBufferSize:  conf.ServerOption.ReadBufferSize,
		WriteBufferSize: conf.ServerOption.WriteBufferSize,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Use(requestid.New())

	return app
}
