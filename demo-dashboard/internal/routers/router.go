package routers

import (
	"demo-dashboard/internal/conf"
	"demo-dashboard/internal/filter"
	"demo-dashboard/internal/handler/route"
	"demo-dashboard/internal/handler/service"
	"demo-dashboard/internal/handler/upstream"
	"demo-dashboard/internal/handler/user"

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

	app.Use(filter.JwtAuthentication())
	// monitor
	app.Get("/metrics", monitor.New())

	/*
		app.Use(jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte("secret")},
		}))
	*/

	api_admin := app.Group("/apisix/admin")
	{
		api_admin.Post("/user/signin", user.SignIn)
		api_admin.Post("user/signout", user.SignOut)
	}

	{
		api_admin.Get("/routes", route.GetList)
	}
	{
		api_admin.Get("/services", service.GetList)
		api_admin.Get("/services/:id([0-9]+)", service.Get)
		api_admin.Put("/services/:id([0-9]+)", service.Put)
		api_admin.Delete("/services/:id([0-9]+)", service.Delete)
		api_admin.Patch("/services/:id([0-9]+)", service.Patch)
		api_admin.Post("/services", service.Post)
	}
	{
		api_admin.Put("/upstreams/:id([0-9]+)", upstream.Put)
		api_admin.Get("/upstreams/:id([0-9]+)", upstream.Get)
		api_admin.Delete("/upstreams/:id([0-9]+)", upstream.Delete)
		api_admin.Patch("/upstreams/:id([0-9]+)", upstream.Patch)
		api_admin.Post("/upstreams", upstream.Post)
		api_admin.Get("/upstreams", upstream.GetList)
		api_admin.Get("/upstreams/:name", upstream.Search)
	}

	return app
}
