package middleware

import (
	"errors"
	UserModel "server/server/internal/models"

	"github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Request().Cookie("Auth")
		if err != nil || cookie != nil && cookie.MaxAge == -1 {
			return c.Redirect(303, "/auth/login")
		} else {
			return next(c)
		}
	}
}

func IsNotAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Request().Cookie("Auth")
		if err == nil || cookie != nil && cookie.MaxAge > -1 {
			return c.Redirect(303, "/auth/logout")
		} else {
			return next(c)
		}
	}
}

func TakesUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("Email")
		if email == "" {
			return errors.New("Email is required")
		}

		username := c.FormValue("Username")
		if username == "" {
			return errors.New("Username is required")
		}

		password := c.FormValue("Password")
		if password == "" {
			return errors.New("Password is required")
		}

		confirm := c.FormValue("Confirm Password")
		if confirm == "" {
			return errors.New("Password confirmation is required")
		}

		if password != confirm {
			return errors.New("Passwords must match")
		}

		user := &UserModel.User{
			Email:    email,
			Username: username,
			Password: password,
		}

		c.Set("User", user)

		return next(c)
	}
}
