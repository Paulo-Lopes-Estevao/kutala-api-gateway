package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/injection"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/api/handler"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/api/middleware"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/database"
	"github.com/labstack/echo"
	echo_middleware "github.com/labstack/echo/middleware"
)

var public_server = echo.New()

var private_server = echo.New()

func main() {

	db := database.ConnectionDB()
	i := injection.NewRegistry(db)

	public_server.Use(echo_middleware.Logger())
	public_server.Use(echo_middleware.Recover())

	middleware.RegistryMiddlewareBasicAuth(private_server, i.NewAppController())

	handler.RouteMicroserviceHandler(private_server, i.NewAppController())

	handler.RouteUserHandler(public_server, i.NewAppController())

	go func() {
		fmt.Println("Running Server A")
		fmt.Println("Server started at port 8080")
		if err := private_server.Start(":8080"); err != http.ErrServerClosed {
			log.Println("Not Running Server A...", err.Error())
		}
	}()

	fmt.Println("Running Server B")
	fmt.Println("Server started at port 8081")
	err := public_server.Start(":8081")

	if err != nil {
		log.Println("Not Running Server B...", err.Error())
	}
}
