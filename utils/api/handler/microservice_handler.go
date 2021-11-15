package handler

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/interface/controller"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/api/middleware"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/api/proxy"
	"github.com/labstack/echo"
)

func RouteMicroserviceHandler(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.POST("/microservice", func(context echo.Context) error { return c.Microservice.AddMicroservice(context) })
	e.POST("/microservice/:id/:name", func(context echo.Context) error { return proxy.ReverseProxy(context, c) })

	middleware.RegistryMiddlewareBasicAuth(e, c)

	return e

}
