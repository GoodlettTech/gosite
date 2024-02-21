package auth

import (
	"fmt"
	"net/http"
	"server/server/internal/middleware"
	UserModel "server/server/internal/models"
	UserService "server/server/internal/services/User"
	"server/server/web/partials"
	Views "server/server/web/views"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func RegisterPages(router *echo.Group) {
	router.GET("/login", func(c echo.Context) error {
		hx := strings.ToLower(c.Request().Header.Get("HX-Request"))

		if hx == "true" {
			return partials.LoginForm().Render(c.Request().Context(), c.Response().Writer)
		} else {
			return Views.Login().Render(c.Request().Context(), c.Response().Writer)
		}
	}, middleware.IsNotAuthenticated)

	router.POST("/login", func(c echo.Context) error {
		cookie := http.Cookie{
			Name:     "Auth",
			Value:    "Test",
			Path:     "/",
			Domain:   "http://localhost:3000",
			Expires:  time.Now().Add(1 * time.Hour),
			Secure:   true,
			HttpOnly: true,
		}
		c.SetCookie(&cookie)

		return c.Redirect(303, "/")
	})

	router.GET("/createaccount", func(c echo.Context) error {
		hx := strings.ToLower(c.Request().Header.Get("HX-Request"))

		if hx == "true" {
			return partials.CreateAccountForm().Render(c.Request().Context(), c.Response().Writer)
		} else {
			return Views.CreateAccount().Render(c.Request().Context(), c.Response().Writer)
		}
	})

	router.POST("/createaccount", func(c echo.Context) error {
		user := c.Get("User").(*UserModel.User)
		UserService.AddUser(user)

		cookie := http.Cookie{
			Name:     "Auth",
			Value:    fmt.Sprintf("%s:%s", user.Email, user.Username),
			Path:     "/",
			Domain:   "http://localhost:3000",
			Expires:  time.Now().Add(1 * time.Hour),
			MaxAge:   int(time.Hour),
			Secure:   true,
			HttpOnly: true,
		}
		c.SetCookie(&cookie)

		return c.Redirect(303, "/")
	}, middleware.TakesUser)

	router.GET("/logout", func(c echo.Context) error {
		_, err := c.Request().Cookie("Auth")
		if err == nil {
			cookie := http.Cookie{
				Name:     "Auth",
				Value:    "",
				Path:     "/",
				Domain:   "http://localhost:3000",
				Expires:  time.Now(),
				MaxAge:   -1,
				Secure:   true,
				HttpOnly: true,
			}

			c.SetCookie(&cookie)
		}
		return c.Redirect(303, "/auth/login")
	})
}
