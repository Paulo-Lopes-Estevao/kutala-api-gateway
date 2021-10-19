package handler

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/controller"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/api/proxy"
	"github.com/labstack/echo"
)

func RouteMicroserviceHandler(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.POST("/microservice", func(context echo.Context) error { return c.Microservice.AddMicroservice(context) })
	e.GET("/v1/africa", echo.WrapHandler(proxy.ReverseProxy()))
	//e.Any("/:pathmicroservice", echo.WrapHandler()))
	return e

}
