package routers

import (
	"demo-dashboard/internal/conf"
	"demo-dashboard/internal/handler/route"
	"demo-dashboard/internal/handler/service"
	"demo-dashboard/internal/handler/upstream"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	logger, _ := zap.NewProduction()

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Use(requestid.New())
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
		Fields: []string{"requestid", "time", "pid", "status", "method", "path", "latencry", "url"},
		Levels: []zapcore.Level{zapcore.FatalLevel, zapcore.PanicLevel, zapcore.ErrorLevel, zapcore.WarnLevel, zapcore.InfoLevel, zapcore.DebugLevel},
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
