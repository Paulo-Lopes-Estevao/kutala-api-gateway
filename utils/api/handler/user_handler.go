package handler

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/interface/controller"
	"github.com/gorilla/mux"
)

func RouteServiceHandler(e *mux.Router, c controller.AppController) *mux.Router {

	e.HandleFunc("/User", c.User.AddUser).Methods("POST")
	e.HandleFunc("/User/{id}", c.User.FindByUuidUser).Methods("GET")

	return e

}
