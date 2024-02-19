package routes

import (
	test "server/server/internal/routes/api"
	views "server/server/internal/routes/pages"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	//register page routes
	registerRoutes(e.Group("/"), views.Routes)

	//register routes to /api
	api := e.Group("/api")
	registerRoutes(api.Group("/test"), test.Routes)
}

type RegisterRouterFunc func(*echo.Group)

func registerRoutes(group *echo.Group, register RegisterRouterFunc) {
	register(group)
}