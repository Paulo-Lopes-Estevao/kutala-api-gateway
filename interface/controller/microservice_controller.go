package controller

import (
	"errors"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/domain/entities"
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/interface/presenter"
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/usecase/interactor"
)

type MicroserviceControllerInterface interface {
	AddMicroservice(ctx Context) error
	GetMicroservice(ctx Context) (*entities.Microservice, error)
}

type microserviceController struct {
	microserviceUseCase interactor.UserUseCaseInterface
}

func NewMicroserviceController(usecases interactor.UserUseCaseInterface) MicroserviceControllerInterface {
	return &microserviceController{usecases}
}

var microservice entities.Microservice

func (usecasecontroller *microserviceController) AddMicroservice(ctx Context) error {

	if err := ctx.Bind(&microservice); !errors.Is(err, nil) {
		ctx.JSON(http.StatusBadRequest, ResponseData{"error": err.Error()})
	}

	microservices, err := usecasecontroller.microserviceUseCase.CreateMicroService(&microservice)

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, ResponseData{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, ResponseData{"data": microservices})

}

func (usecasecontroller *microserviceController) GetMicroservice(ctx Context) (*entities.Microservice, error) {

	id, pathmicroservice := ctx.Param("id"), ctx.Param("name")

	users, err := usecasecontroller.microserviceUseCase.SearchUuid(id, &user)

	if err != nil {
		return nil, err
	}

	microservices, err := usecasecontroller.microserviceUseCase.SearchPathService(pathmicroservice, &microservice)

	if err != nil {
		return nil, err
	}

	return presenter.Response(microservices, users), nil

}
