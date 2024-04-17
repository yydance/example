package routers

import (
	"demo-dashboard/internal/conf"
	"demo-dashboard/internal/handler/route"
	"demo-dashboard/internal/handler/service"
	"demo-dashboard/internal/handler/upstream"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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
	app.Use(logger.New(logger.Config{
		Format:     "[${locals:requestid}] ${time} ip: ${ip} status: ${status} pid: ${pid} method: ${method} path: ${path} queryParams: ${queryParams} body: ${body} resBody: ${resBody}\n error: ${error}\n",
		TimeFormat: time.RFC3339Nano,
		TimeZone:   "Asia/Shanghai",
	}))
	// monitor
	app.Get("/metrics", monitor.New())

	/*
		app.Use(jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte("secret")},
		}))
	*/

	api_admin := app.Group("/apisix/admin")
	{
		api_admin.Put("/upstreams/:id", upstream.Put)
		api_admin.Get("/upstreams/:id", upstream.Get)
		api_admin.Delete("/upstreams/:id", upstream.Delete)
		api_admin.Patch("/upstreams/:id", upstream.Patch)
		api_admin.Post("/upstreams", upstream.Post)
	}
	{
		api_admin.Get("/services", service.GetList)
		api_admin.Get("/services/:id", service.Get)
		api_admin.Put("/services/:id", service.Put)
		api_admin.Delete("/services/:id", service.Delete)
		api_admin.Patch("/services/:id", service.Patch)
		api_admin.Post("/services", service.Post)
	}

	{
		api_admin.Get("/routes", route.GetList)
	}

	return app
}
