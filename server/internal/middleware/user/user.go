package UserMiddleware

import (
	"errors"
	UserModel "server/server/internal/models"

	"github.com/labstack/echo/v4"
)

func TakesUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("Email")
		if email == "" {
			return errors.New("email is required")
		}

		username := c.FormValue("Username")
		if username == "" {
			return errors.New("username is required")
		}

		password := c.FormValue("Password")
		if password == "" {
			return errors.New("password is required")
		}

		confirm := c.FormValue("Confirm Password")
		if confirm == "" {
			return errors.New("password confirmation is required")
		}

		if password != confirm {
			return errors.New("passwords must match")
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

func TakesCredentials(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("Username")
		if username == "" {
			return errors.New("username is required")
		}

		password := c.FormValue("Password")
		if password == "" {
			return errors.New("password is required")
		}

		creds := &UserModel.Credentials{
			Username: username,
			Password: password,
		}

		c.Set("credentials", creds)

		return next(c)
	}
}
