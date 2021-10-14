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
)

var private_server = echo.New()
var public_server = echo.New()

func main() {
	db := database.ConnectionDB()
	i := injection.NewRegistry(db)

	handler.RouteServiceHandler(public_server, i.NewAppController())

	middleware.RegistryMiddleware(private_server, i.NewAppController())
	handler.RouteMicroserviceHandler(private_server, i.NewAppController())

	go func() {
		fmt.Println("Running Server A")
		if err := private_server.Start(":8080"); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	fmt.Println("Running Server B")
	if err := public_server.Start(":8081"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
