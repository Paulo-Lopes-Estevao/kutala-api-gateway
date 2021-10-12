package handler

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/controller"
	"github.com/labstack/echo"
)

func PrivateRouteMicroserviceHandler(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.GET("/microservice/:id", func(context echo.Context) error { return c.Microservice.FindByUuidMicroservice(context) })

	return e

}


func PublicRouteMicroserviceHandler(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.POST("/microservice", func(context echo.Context) error { return c.Microservice.AddMicroservice(context) })

	return e

}