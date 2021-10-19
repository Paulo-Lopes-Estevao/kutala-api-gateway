package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/usecase"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/gorilla/mux"
)

type ServiceControllerInterface interface {
	AddService(w http.ResponseWriter, r *http.Request)
	AuthBasicService(username, password string) (*entities.Service, error)
	FindByUuidService(w http.ResponseWriter, r *http.Request)
}

type serviceController struct {
	serviceUseCase usecase.ServiceUseCaseInterface
}

func NewServiceController(usecases usecase.ServiceUseCaseInterface) ServiceControllerInterface {
	return &serviceController{usecases}
}

var service entities.Service

func (usecasecontroller *serviceController) AddService(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&service)

	services, err := usecasecontroller.serviceUseCase.CreateService(&service)

	if err != nil {
		fmt.Fprint(w, err)
	}

	resp := make(map[string]interface{})
	resp["message"] = "Status Created"
	resp["data"] = services
	resp["status"] = http.StatusCreated

	value, _ := json.Marshal(resp)

	w.Write(value)
}

func (usecasecontroller *serviceController) AuthBasicService(username, password string) (*entities.Service, error) {

	services, err := usecasecontroller.serviceUseCase.Auth(username, password, &service)

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (usecasecontroller *serviceController) FindByUuidService(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	services, err := usecasecontroller.serviceUseCase.SearchUuid(id["id"], &service)

	if !errors.Is(err, nil) {
		fmt.Fprint(w, err)
	}

	resp := make(map[string]interface{})
	resp["message"] = "Service Found"
	resp["data"] = services
	resp["status"] = http.StatusOK

	value, _ := json.Marshal(resp)

	w.Write(value)

}
