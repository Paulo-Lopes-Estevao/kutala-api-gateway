package main

import (
	"fmt"
	"log"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/injection"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/api/handler"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var server = echo.New()

func main() {

	db := database.ConnectionDB()
	i := injection.NewRegistry(db)

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	handler.RouteMicroserviceHandler(server, i.NewAppController())

	handler.RouteUserHandler(server, i.NewAppController())

	fmt.Println("Server started at port 9000")
	err := server.Start(":9000")

	if err != nil {
		log.Println("Not Running Server ...", err.Error())
	}
}
