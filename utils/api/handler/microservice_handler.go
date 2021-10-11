package handler

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RegistryRoute(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/test", func(context echo.Context) error { return c.Microservice.AddMicroservice(context) })

	return e

}
