package usecase

import (
	"fmt"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/app/repository"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
)

type UserUseCaseInterface interface {
	CreateUser(user *entities.User) (*entities.User, error)
	Auth(username string, password string, user *entities.User) (*entities.User, error)
	SearchUuid(uuid string, user *entities.User) (*entities.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserUseCase(repository repository.UserRepositoryInterface) UserUseCaseInterface {
	return &userUseCase{repository}
}

func (usecase *userUseCase) CreateUser(user *entities.User) (*entities.User, error) {

	data, err := entities.NewUser(user.Name, user.Username, user.Password)

	if err != nil {
		return nil, err
	}

	services, err := usecase.userRepository.InsertUser(data)

	if err != nil {
		return nil, err
	}

	return services, nil
}

func (usecase *userUseCase) Auth(username string, password string, user *entities.User) (*entities.User, error) {
	data, err := usecase.userRepository.FindUsernameService(username, user)

	if err != nil {
		return nil, fmt.Errorf("The password is invalid for the username")
	}

	if data.VerifyPassword(password) {
		return data, nil
	}

	return nil, fmt.Errorf("The password is invalid for the username")
}

func (usecase *userUseCase) SearchUuid(uuid string, user *entities.User) (*entities.User, error) {
	data, err := usecase.userRepository.FindIdUser(uuid, user)

	if err != nil {
		return nil, fmt.Errorf("Id is invalid")
	}

	return data, nil
}
