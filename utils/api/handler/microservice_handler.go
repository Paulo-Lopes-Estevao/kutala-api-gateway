package handler

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/interface/controller"
	"github.com/labstack/echo"
)

func RouteMicroserviceHandler(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.POST("/microservice/:id", func(context echo.Context) error { return c.Microservice.AddMicroservice(context) })

	return e

}
