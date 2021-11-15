package controller

import (
	"errors"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/usecase/interactor"
)

type UserControllerInterface interface {
	AddUser(ctx Context) error
	AuthBasicUser(username, password string, ctx Context) (*entities.User, error)
	FindByUuidUser(ctx Context) error
}

type userController struct {
	userUseCase interactor.UserUseCaseInterface
}

func NewUserController(usecases interactor.UserUseCaseInterface) UserControllerInterface {
	return &userController{usecases}
}

var user entities.User

func (usecasecontroller *userController) AddUser(ctx Context) error {

	if err := ctx.Bind(&user); !errors.Is(err, nil) {
		ctx.JSON(http.StatusBadRequest, ResponseData{"error": err.Error()})
	}

	users, err := usecasecontroller.userUseCase.CreateUser(&user)

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, ResponseData{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, ResponseData{"data": users})
}

func (usecasecontroller *userController) AuthBasicUser(username, password string, ctx Context) (*entities.User, error) {

	services, err := usecasecontroller.userUseCase.Auth(username, password, &user)

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (usecasecontroller *userController) FindByUuidUser(ctx Context) error {
	id := ctx.Param("id")

	users, err := usecasecontroller.userUseCase.SearchUuid(id, &user)

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusNotFound, ResponseData{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, ResponseData{"data": users})

}
