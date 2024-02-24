package AuthMiddleware

import (
	"fmt"
	"net/url"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Request().Cookie("Auth")

		redirectUrl := "/auth/login"

		//hx request coming from the page
		if c.Request().Header.Get("HX-Request") == "true" {
			hxCurrent := c.Request().Header.Get("HX-Current-URL")
			redirectUrl = fmt.Sprintf("%s?redirect=%s", redirectUrl, url.QueryEscape(hxCurrent))
		} else //browser request loading the page for the first time
		if currentUrl := c.Request().URL.String(); currentUrl != "/auth/login" {
			redirectUrl = fmt.Sprintf("%s?redirect=%s", redirectUrl, url.QueryEscape(currentUrl))
		}

		if err != nil || cookie != nil && cookie.MaxAge == -1 {
			return c.Redirect(303, redirectUrl)
		} else {
			token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if token.Valid && err == nil {
				return next(c)
			} else {
				return c.Redirect(303, redirectUrl)
			}
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
