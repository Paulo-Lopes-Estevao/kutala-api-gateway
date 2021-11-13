package handler

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/interface/controller"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/utils/api/proxy"
	"github.com/gorilla/mux"
)

func RouteMicroserviceHandler(e *mux.Router, c controller.AppController) *mux.Router {

	e.HandleFunc("/microservice", c.Microservice.AddMicroservice).Methods("POST")
	e.HandleFunc("/graphql", proxy.Handler(proxy.ReverseProxy())).Methods("GET")

	return e

}
