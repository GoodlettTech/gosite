package auth

import (
	"errors"
	"fmt"
	"net/url"
	AuthMiddleware "server/server/internal/middleware/auth"
	UserMiddleware "server/server/internal/middleware/user"
	UserModel "server/server/internal/models"
	AuthService "server/server/internal/services/auth"
	UserService "server/server/internal/services/user"
	"server/server/web/partials"
	Views "server/server/web/views"
	"strings"

	"github.com/labstack/echo/v4"
)

func RegisterPages(router *echo.Group) {
	router.GET("/login", func(c echo.Context) error {
		hx := strings.ToLower(c.Request().Header.Get("HX-Request"))

		redirect := c.QueryParams().Get("redirect")
		url, err := url.QueryUnescape(redirect)
		if err != nil {
			panic(err)
		}
		fmt.Println(url)

		if hx == "true" {
			return partials.LoginForm().Render(c.Request().Context(), c.Response().Writer)
		} else {
			return Views.Login().Render(c.Request().Context(), c.Response().Writer)
		}
	}, AuthMiddleware.IsNotAuthenticated)

	router.POST("/login", func(c echo.Context) error {
		creds := c.Get("credentials").(*UserModel.Credentials)
		userId, err := UserService.VerifyUser(creds)
		if err != nil {
			return err
		}
		if userId == -1 {
			return errors.New("failed to verify user")
		}

		cookie := AuthService.CreateCookie(userId)
		c.SetCookie(&cookie)

		return c.Redirect(303, "/")
	}, UserMiddleware.TakesCredentials, AuthMiddleware.IsNotAuthenticated)

	router.GET("/createaccount", func(c echo.Context) error {
		hx := strings.ToLower(c.Request().Header.Get("HX-Request"))

		if hx == "true" {
			return partials.CreateAccountForm().Render(c.Request().Context(), c.Response().Writer)
		} else {
			return Views.CreateAccount().Render(c.Request().Context(), c.Response().Writer)
		}
	}, AuthMiddleware.IsNotAuthenticated)

	router.POST("/createaccount", func(c echo.Context) error {
		user := c.Get("User").(*UserModel.User)
		user.Id = -1

		UserService.AddUser(user)
		if user.Id == -1 {
			return errors.New("user creation failed")
		}

		cookie := AuthService.CreateCookie(user.Id)
		c.SetCookie(&cookie)

		return c.Redirect(303, "/")
	}, AuthMiddleware.IsNotAuthenticated, UserMiddleware.TakesUser)

	router.GET("/logout", func(c echo.Context) error {
		cookie := AuthService.CreateEmptyCookie()
		c.SetCookie(&cookie)
		return c.Redirect(303, "/auth/login")
	})
}
