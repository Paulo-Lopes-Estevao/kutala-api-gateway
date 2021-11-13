package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/usecase/interactor"
	"github.com/gorilla/mux"
)

type UserControllerInterface interface {
	AddUser(w http.ResponseWriter, r *http.Request)
	AuthBasicUser(username, password string) (*entities.User, error)
	FindByUuidUser(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	serviceUseCase interactor.UserUseCaseInterface
}

func NewUserController(usecases interactor.UserUseCaseInterface) UserControllerInterface {
	return &userController{usecases}
}

var User entities.User

func (usecasecontroller *userController) AddUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&User)

	services, err := usecasecontroller.serviceUseCase.CreateUser(&User)

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

func (usecasecontroller *userController) AuthBasicUser(username, password string) (*entities.User, error) {

	services, err := usecasecontroller.serviceUseCase.Auth(username, password, &User)

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (usecasecontroller *userController) FindByUuidUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	services, err := usecasecontroller.serviceUseCase.SearchUuid(id["id"], &User)

	if !errors.Is(err, nil) {
		fmt.Fprint(w, err)
	}

	resp := make(map[string]interface{})
	resp["message"] = "User Found"
	resp["data"] = services
	resp["status"] = http.StatusOK

	value, _ := json.Marshal(resp)

	w.Write(value)

}
