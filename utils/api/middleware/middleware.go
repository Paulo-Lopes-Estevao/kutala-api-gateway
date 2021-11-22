package middleware

import (
	"crypto/subtle"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/interface/controller"
	"github.com/labstack/echo/v4"
)

func RegistryMiddlewareBasicAuth(e *echo.Echo, c controller.AppController) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {

			username, password, ok := context.Request().BasicAuth()

			if !ok {
				return context.JSON(401, "Unauthorized.")
			}

			users, err := c.User.AuthBasicUser(username, password, context)

			if err != nil {
				return context.JSON(401, "Unauthorized.")
			}

			result := users.VerifyPassword(password)

			if subtle.ConstantTimeCompare([]byte(username), []byte(users.Username)) != 1 && !result {
				return context.JSON(401, "Unauthorized.")
			}
			return next(context)
		}
	}
}
