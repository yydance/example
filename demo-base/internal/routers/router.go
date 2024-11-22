package routers

import (
	"demo-base/internal/conf"
	"demo-base/internal/handler"
	"demo-base/internal/middleware/jwt"
	"demo-base/internal/middleware/opa"
	"demo-base/internal/utils/logger"
	"time"

	"demo-base/internal/handler/project"
	"demo-base/internal/handler/role"
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
		ServerHeader:          "Olio-Server",
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
		Logger: logger.NewCustomLogger(logger.AccessLog),
	}))
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessEndpoint:  "/live",
		ReadinessEndpoint: "/ready",
	}))
	app.Get("/metrics", monitor.New())
	app.Use(jwt.Authentication())

	// Routes
	routers := app.Group("/api/v1")
	routers.Post("/login", handler.Login)
	/*
		{
			routers.Get("/ping", system.Ping)
			routers.Get("/health", system.Health)
			routers.Get("/info", system.Info)
		}
	*/
	routers.Use(opa.OPA())
	userRouters := routers.Group("/users")
	{
		userRouters.Get("", user.List)
		userRouters.Post("", user.Create)
		userRouters.Get("/:username", user.Get)
		userRouters.Put("/:username", user.Update)
		userRouters.Put("/:username/password", user.UpdatePassword)
		userRouters.Delete("/:username", user.Delete)
	}
	roleRouters := routers.Group("/roles")
	{
		//roleRouters.Get("/list", role.GetAll)
		roleRouters.Post("", role.Create)
		//roleRouters.Get("/detail/:id", user.Get)
		roleRouters.Put("/:roleplatform", role.Update)
		//roleRouters.Delete("/delete/:id", user.Delete)
	}

	projectRouters := routers.Group("/projects")
	{
		projectRouters.Get("", project.List)
		projectRouters.Post("", project.Create)
		//projectRouters.Get("/detail/:name", project.Get)
		projectRouters.Put("/:project", project.Update)
		projectRouters.Delete("/:project", project.Delete)
		projectRouters.Post("/:project/roles", project.CreateRole)
		projectRouters.Delete("/:project/roles/:roleproject", project.DeleteRole)
		projectRouters.Put("/:project/roles/:roleproject", project.UpdateRole)
		projectRouters.Post("/:project/members", project.AddMember)
		projectRouters.Put("/:project/members/:member", project.UpdateMember)
		projectRouters.Delete("/:project/members/:member", project.RemoveMember)
	}
	/*
		appRouters := routers.Group("/apps")
		{
			appRouters.Get("", app.List)
			appRouters.Post("", app.Create)
			appRouters.Put("/:appName", app.Update)
			appRouters.Delete("/:appName", app.Delete)
		}
	*/
	return app
}
