package main

import (
	"fmt"
	"log"

	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/injection"

	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/utils/api/handler"
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/utils/database"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func main() {

	server := echo.New()

	db := database.ConnectionDB()
	i := injection.NewRegistry(db)

	//server.Use(middleware.Logger())
	//server.Use(middleware.Recover())

	p := prometheus.NewPrometheus("echo", nil)

	handler.RouteMicroserviceHandler(server, i.NewAppController())

	handler.RouteUserHandler(server, i.NewAppController())

	p.Use(server)

	fmt.Println("Server started at port 9000")
	if err := server.Start(":9000"); err != nil {
		log.Println("Not Running Server A...", err.Error())
	}
}
