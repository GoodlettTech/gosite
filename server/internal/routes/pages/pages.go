package pages

import (
	"server/server/internal/routes/pages/auth"
	"server/server/internal/routes/pages/home"

	"github.com/labstack/echo/v4"
)

func RegisterPages(router *echo.Group) {
	home.RegisterPages(router.Group(""))

	authRouter := router.Group("auth")

	auth.RegisterPages(authRouter)
}
