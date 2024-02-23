package home

import (
	AuthMiddleware "server/server/internal/middleware/auth"
	Views "server/server/web/views"

	"github.com/labstack/echo/v4"
)

func RegisterPages(router *echo.Group) {
	router.GET("", func(c echo.Context) error {
		return Views.Home().Render(c.Request().Context(), c.Response().Writer)
	}, AuthMiddleware.IsAuthenticated)
}
