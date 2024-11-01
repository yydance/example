package routers

import (
	"demo-base/internal/conf"
	"demo-base/internal/handler"
	"demo-base/internal/middleware/casbin"
	"demo-base/internal/middleware/jwt"
	"demo-base/internal/utils/logger"
	"time"

	"demo-base/internal/handler/project"
	"demo-base/internal/handler/role"
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
		ServerHeader:          "Oreo-Server",
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

	app.Use(jwt.Authentication())

	// Routes
	app.Post("/login", handler.Login)
	routers := app.Group("/api/v1")
	{
		routers.Get("/ping", system.Ping)
		routers.Get("/health", system.Health)
		routers.Get("/info", system.Info)
	}
	userRouters := routers.Group("/users")
	{
		userRouters.Get("/list", casbin.RoutePermission(), user.List)
		userRouters.Post("/create", casbin.RoutePermission(), user.Create)
		userRouters.Get("/:name", casbin.RoutePermission(), user.Get)
		userRouters.Put("/:name", casbin.RoutePermission(), user.Update)
		userRouters.Put("/:name/password", casbin.RoutePermission(), user.UpdatePassword)
		userRouters.Delete("/:name", casbin.RoutePermission(), user.Delete)
	}
	roleRouters := routers.Group("/roles")
	{
		//roleRouters.Get("/list", role.GetAll)
		roleRouters.Post("/create", casbin.RoutePermission(), role.Create)
		//roleRouters.Get("/detail/:id", user.Get)
		roleRouters.Put("/:name", casbin.RoutePermission(), role.Update)
		//roleRouters.Delete("/delete/:id", user.Delete)
	}

	projectRouters := routers.Group("/projects")
	{
		projectRouters.Get("/list", casbin.RoutePermission(), project.List)
		projectRouters.Post("/create", casbin.RoutePermission(), project.Create)
		//projectRouters.Get("/detail/:name", project.Get)
		projectRouters.Put("/:name", casbin.RoutePermission(), project.Update)
		projectRouters.Delete("/:name", project.Delete)
		projectRouters.Post("/:name/roles", casbin.RoutePermission(), project.CreateRole)
		projectRouters.Delete("/:name/roles/:roleName", casbin.RoutePermission(), project.DeleteRole)
		projectRouters.Put("/:name/roles/:roleName", casbin.RoutePermission(), project.UpdateRole)
		projectRouters.Post("/:name/members", casbin.RoutePermission(), project.AddMember)
		projectRouters.Put("/:name/members/:memberName", casbin.RoutePermission(), project.UpdateMember)
		projectRouters.Delete("/:name/members/:memberName", casbin.RoutePermission(), project.RemoveMember)
	}

	return app
}
