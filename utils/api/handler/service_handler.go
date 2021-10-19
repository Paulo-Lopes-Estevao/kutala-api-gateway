package handler

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/controller"
	"github.com/gorilla/mux"
)

func RouteServiceHandler(e *mux.Router, c controller.AppController) *mux.Router {

	e.HandleFunc("/service", c.Service.AddService).Methods("POST")
	e.HandleFunc("/service/{id}", c.Service.FindByUuidService).Methods("GET")

	return e

}
