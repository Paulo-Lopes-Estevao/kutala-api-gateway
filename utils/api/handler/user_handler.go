package handler

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/interface/controller"
	"github.com/labstack/echo/v4"
)

func RouteUserHandler(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.POST("/user", func(context echo.Context) error { return c.User.AddUser(context) })
	e.GET("/user/:id", func(context echo.Context) error { return c.User.FindByUuidUser(context) })

	return e

}
