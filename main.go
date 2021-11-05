package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/injection"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/api/handler"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/database"
	"github.com/gorilla/mux"
)

var public_server = mux.NewRouter()

//var e = echo.New()

func main() {

	db := database.ConnectionDB()
	i := injection.NewRegistry(db)

	handler.RouteMicroserviceHandler(public_server, i.NewAppController())

	handler.RouteServiceHandler(public_server, i.NewAppController())

	//middleware.RegistryMiddleware(private_server, i.NewAppController())
	//	handler.RouteMicroserviceHandler(private_server, i.NewAppController())

	/* 	go func() {
	   		fmt.Println("Running Server A")
	   		if err := private_server.Start(":8080"); err != http.ErrServerClosed {
	   			log.Fatal(err)
	   		}
	   	}()
	*/
	fmt.Println("Running Server B")
	fmt.Println("Server started at port 9990")
	err := http.ListenAndServe(":9990", public_server)

	if err != nil {
		log.Println("Not Running Server...", err.Error())
	}
}
