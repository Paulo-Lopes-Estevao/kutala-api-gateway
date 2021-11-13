package middleware

import (
	"crypto/subtle"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RegistryMiddlewareBasicAuth(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.BasicAuth(func(username, password string, context echo.Context) (bool, error) {

		users, _ := c.User.AuthBasicUser(username, password, context)

		result := users.VerifyPassword(password)

		if subtle.ConstantTimeCompare([]byte(username), []byte(users.Username)) == 1 && result {
			return true, nil
		}
		return false, nil

	}))

	return e

}
