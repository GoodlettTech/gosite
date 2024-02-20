package views

import (
	Home "server/server/web/views"

	"github.com/labstack/echo/v4"
)

func Routes(group *echo.Group) {
	group.GET("", func(c echo.Context) error {
		return Home.Home().Render(c.Request().Context(), c.Response().Writer)
	})
}
