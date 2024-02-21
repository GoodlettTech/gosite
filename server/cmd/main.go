package main

import (
	"server/server/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	routes.InitRoutes(e)

	e.Logger.Fatal(
		e.Start(":3000"))
}
