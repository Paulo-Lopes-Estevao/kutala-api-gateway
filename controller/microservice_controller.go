package controller

import (
	"errors"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/usecase"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
)

type MicroserviceControllerInterface interface {
	AddMicroservice(c Context) error
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
