package controller

import (
	"errors"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/usecase"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
)

type MicroserviceControllerInterface interface {
	AddMicroservice(c Context) error
	//ProxyMicroservice(ctx Context) error
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

func verifyPathMicroService(pathmicroservice string, usecasecontroller *microserviceController) (string, error) {

	path, _ := target(pathmicroservice)

	_, err := usecasecontroller.microserviceUseCase.SearchPathService(path, &microservice)

	if err != nil {
		return path, err
	}

	return path, nil
}

/* func (usecasecontroller *microserviceController) ProxyMicroservice(path string) (string, error) {

	pathmicroservice, _ := verifyPathMicroService(path, usecasecontroller)

	microservices, err := usecasecontroller.microserviceUseCase.SearchPathService(pathmicroservice, &microservice)

	if err != nil {
		return nil, err
	}

	return microservices.Api, nil

} */
