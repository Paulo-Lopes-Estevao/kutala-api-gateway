package controller

import (
	"errors"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/usecase"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
)

type ServiceControllerInterface interface {
	AddService(c Context) error
	AuthBasicService(username, password string, c Context) (*entities.Service, error)
	FindByUuidService(c Context) error
}

type serviceController struct {
	serviceUseCase usecase.ServiceUseCaseInterface
}

func NewServiceController(usecases usecase.ServiceUseCaseInterface) ServiceControllerInterface {
	return &serviceController{usecases}
}

var service entities.Service

func (usecasecontroller *serviceController) AddService(ctx Context) error {

	if err := ctx.Bind(&service); !errors.Is(err, nil) {

		return ctx.JSON(http.StatusBadRequest, ResponseData{"error": err.Error()})
	}

	services, err := usecasecontroller.serviceUseCase.CreateService(&service)

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, ResponseData{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, ResponseData{"data": services})
}

func (usecasecontroller *serviceController) AuthBasicService(username, password string, ctx Context) (*entities.Service, error) {

	services, err := usecasecontroller.serviceUseCase.Auth(username, password, &service)

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (usecasecontroller *serviceController) FindByUuidService(ctx Context) error {
	id := ctx.Param("id")

	services, err := usecasecontroller.serviceUseCase.SearchUuid(id, &service)

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusNotFound, ResponseData{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, ResponseData{"data": services})
}
