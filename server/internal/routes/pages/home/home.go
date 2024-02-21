package home

import (
	"server/server/internal/middleware"
	Views "server/server/web/views"

	"github.com/labstack/echo/v4"
)

func RegisterPages(router *echo.Group) {
	router.GET("", func(c echo.Context) error {
		return Views.Home().Render(c.Request().Context(), c.Response().Writer)
	}, middleware.IsAuthenticated)
}
