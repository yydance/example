package routers

import (
	"demo-base/internal/conf"
	"demo-base/internal/utils/logger"
	"time"

	"demo-base/internal/handler/system"
	"demo-base/internal/handler/user"

	"github.com/bytedance/sonic"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               conf.FiberConfig.AppName,
		ReadTimeout:           conf.FiberConfig.ReadTimeout * time.Microsecond,
		WriteTimeout:          conf.FiberConfig.WriteTimeout * time.Microsecond,
		ReadBufferSize:        conf.FiberConfig.ReadBufferSize,
		WriteBufferSize:       conf.FiberConfig.WriteBufferSize,
		Concurrency:           conf.FiberConfig.Concurrent,
		Prefork:               conf.FiberConfig.Prefork,
		IdleTimeout:           conf.FiberConfig.IdleTimeout * time.Microsecond,
		Network:               conf.FiberConfig.Network,
		BodyLimit:             conf.FiberConfig.BodyLimit,
		JSONEncoder:           sonic.Marshal,
		JSONDecoder:           sonic.Unmarshal,
		DisableStartupMessage: false,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	if conf.CorsConfig.Enabled {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     conf.CorsConfig.AllowOrigins,
			AllowMethods:     conf.CorsConfig.AllowMethods,
			AllowHeaders:     conf.CorsConfig.AllowHeaders,
			AllowCredentials: conf.CorsConfig.AllowCredentials,
			ExposeHeaders:    conf.CorsConfig.ExposeHeaders,
			MaxAge:           conf.CorsConfig.MaxAge,
		}))
	}

	app.Use(requestid.New(requestid.Config{}))

	app.Use(fiberzap.New(fiberzap.Config{
		Fields: []string{"ip", "ips", "method", "protocol", "referer", "url", "route", "ua", "status", "latency", "bytesReceived", "bytesSent", "error", "requestId"},
		Levels: logger.LogLevel(),
	}))
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessEndpoint:  "/live",
		ReadinessEndpoint: "/ready",
	}))
	app.Get("/metrics", monitor.New())

	// Routes
	routers := app.Group("/api/v1")
	{
		routers.Get("/ping", system.Ping)
		routers.Get("/health", system.Health)
		routers.Get("/info", system.Info)
	}
	userRouters := routers.Group("/user")
	{
		userRouters.Get("/list", user.GetAll)
		userRouters.Post("/create", user.Create)
		userRouters.Get("/detail/:id", user.Get)
		userRouters.Put("/update/:id", user.Update)
		userRouters.Delete("/delete/:id", user.Delete)
	}

	return app
}
