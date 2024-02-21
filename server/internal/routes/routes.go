package routes

import (
	"server/server/internal/routes/pages"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	RegisterRoutes(e.Group(""), pages.RegisterPages)
}

type RegisterRouterFunc func(*echo.Group)

func RegisterRoutes(group *echo.Group, register RegisterRouterFunc) {
	register(group)
}
