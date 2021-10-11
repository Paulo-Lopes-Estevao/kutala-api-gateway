package middleware

import (
	"crypto/subtle"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RegistryMiddleware(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.BasicAuth(func(username, password string, context echo.Context) (bool, error) {

		microservices, _ := c.Microservice.AuthBasicMicroservice(username, password, context)

		result := microservices.VerifyPassword(password)

		if subtle.ConstantTimeCompare([]byte(username), []byte(microservices.Username)) == 1 && result {
			return true, nil
		}
		return false, nil

	}))

	return e

}
