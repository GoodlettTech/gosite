package test

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(group *echo.Group) {
	group.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
