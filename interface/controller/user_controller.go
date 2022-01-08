package controller

import (
	"errors"
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/bootstrap"
	"net/http"
	"time"

	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/domain/entities"
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/usecase/interactor"
	"github.com/golang-jwt/jwt"
)

type UserControllerInterface interface {
	AddUser(ctx Context) error
	AuthBasicUser(username, password string, ctx Context) (*entities.User, error)
	AuthJWTUser(ctx Context) error
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

type jwtCustomClaims struct {
	IdUser     string `json:"id_user"`
	Authorized bool   `json:"authorized"`
	jwt.StandardClaims
}

func (usecasecontroller *userController) AuthJWTUser(ctx Context) error {

	signingKey := []byte(bootstrap.GoDotEnvVariable("JWT_SECRET_KEY"))

	if err := ctx.Bind(&user); !errors.Is(err, nil) {
		ctx.JSON(http.StatusBadRequest, ResponseData{"error": err.Error()})
	}

	services, err := usecasecontroller.userUseCase.Auth(user.Username, user.Password, &user)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseData{"error": err.Error()})
	}

	claims := &jwtCustomClaims{
		services.Id,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(signingKey)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseData{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, ResponseData{
		"token": t,
	})
}

func (usecasecontroller *userController) FindByUuidUser(ctx Context) error {
	id := ctx.Param("id")

	users, err := usecasecontroller.userUseCase.SearchUuid(id, &user)

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusNotFound, ResponseData{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, ResponseData{"data": users})

}
