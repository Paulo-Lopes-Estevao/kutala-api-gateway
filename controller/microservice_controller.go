package controller

import (
	"errors"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/usecase"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
)

type MicroserviceControllerInterface interface {
	AddMicroservice(c Context) error
	AuthBasicMicroservice(username, password string, c Context) (*entities.Microservice, error)
	FindByUuidMicroservice(c Context) error
}

type microserviceController struct {
	microserviceUseCase usecase.MicroserviceUseCaseInterface
}

func NewMicroserviceController(usecases usecase.MicroserviceUseCaseInterface) MicroserviceControllerInterface {
	return &microserviceController{usecases}
}

var microservice entities.Microservice

func (usecasecontroller *microserviceController) AddMicroservice(ctx Context) error {

	if err := ctx.Bind(&microservice); !errors.Is(err, nil) {

		return ctx.JSON(http.StatusBadRequest, ResponseData{"error": err.Error()})
	}

	microservices, err := usecasecontroller.microserviceUseCase.CreateMicroService(&microservice)

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, ResponseData{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, ResponseData{"data": microservices})
}

func (usecasecontroller *microserviceController) AuthBasicMicroservice(username, password string, ctx Context) (*entities.Microservice, error) {

	microservices, err := usecasecontroller.microserviceUseCase.Auth(username, password, &microservice)

	if err != nil {
		return nil, err
	}

	return microservices, nil
}

func (usecasecontroller *microserviceController) FindByUuidMicroservice(ctx Context) error {
	id := ctx.Param("id")

	microservices, err := usecasecontroller.microserviceUseCase.SearchUuid(id, &microservice)

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusNotFound, ResponseData{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, ResponseData{"data": microservices})
}
