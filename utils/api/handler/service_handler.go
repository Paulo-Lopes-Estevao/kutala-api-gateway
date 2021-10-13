package handler

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/controller"
	"github.com/labstack/echo"
)



func RouteServiceHandler(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.POST("/service", func(context echo.Context) error { return c.Service.AddService(context) })
	e.GET("/service/:id", func(context echo.Context) error { return c.Service.FindByUuidService(context) })

	return e

}
