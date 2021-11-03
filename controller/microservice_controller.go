package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/usecase"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
)

type MicroserviceControllerInterface interface {
	AddMicroservice(w http.ResponseWriter, r *http.Request)
	//ProxyReverseMicroservice(w http.ResponseWriter, r *http.Request) *httputil.ReverseProxy
}

type microserviceController struct {
	microserviceUseCase usecase.MicroserviceUseCaseInterface
}

func NewMicroserviceController(usecases usecase.MicroserviceUseCaseInterface) MicroserviceControllerInterface {
	return &microserviceController{usecases}
}

var microservice entities.Microservice

func target(path string) (string, error) {
	parts := path
	return parts, nil
}

func (usecasecontroller *microserviceController) AddMicroservice(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&microservice)

	microservices, err := usecasecontroller.microserviceUseCase.CreateMicroService(&microservice)

	if err != nil {
		fmt.Fprint(w, err)
	}

	resp := make(map[string]interface{})
	resp["message"] = "Status Created"
	resp["data"] = microservices
	resp["status"] = http.StatusCreated

	value, _ := json.Marshal(resp)

	w.Write(value)

}

/* func verifyPathMicroService(pathmicroservice string, usecasecontroller *microserviceController) (string, error) {

	path, _ := target(pathmicroservice)

	_, err := usecasecontroller.microserviceUseCase.SearchPathService(path, &microservice)

	if err != nil {
		return path, err
	}

	return path, nil
} */

/* func (usecasecontroller *microserviceController) ProxyReverseMicroservice(w http.ResponseWriter, r *http.Request) *httputil.ReverseProxy {

	//microservicePath := mux.Vars(r)

	//w.WriteHeader(http.StatusOK)

	//pathmicroservice, _ := verifyPathMicroService(microservicePath["microservice"], usecasecontroller)

	microservices, err := usecasecontroller.microserviceUseCase.SearchPathService(pathmicroservice, &microservice)

	if err != nil {
		return nil, err
	}

	return microservices.Api, nil

} */
