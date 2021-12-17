package interactor

import (
	"fmt"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain/entities"
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/usecase/repository"
)

type UserUseCaseInterface interface {
	CreateUser(user *entities.User) (*entities.User, error)
	Auth(username string, password string, user *entities.User) (*entities.User, error)
	SearchUuid(uuid string, user *entities.User) (*entities.User, error)
	CreateMicroService(microservice *entities.Microservice) (*entities.Microservice, error)
	SearchPathService(path string, microservice *entities.Microservice) (*entities.Microservice, error)
}

type userUseCase struct {
	userRepository         repository.UserRepositoryInterface
	microserviceRepository repository.MicroserviceRepositoryInterface
}

func NewUserUseCase(repositoryUser repository.UserRepositoryInterface, repositoryMicroservice repository.MicroserviceRepositoryInterface) UserUseCaseInterface {
	return &userUseCase{repositoryUser, repositoryMicroservice}
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

func (usecase *userUseCase) CreateMicroService(microservice *entities.Microservice) (*entities.Microservice, error) {

	data, err := entities.NewMicroservice(microservice.Name, microservice.Version, microservice.Api, microservice.Endpoint, microservice.Type, microservice.Method, microservice.Iduser)

	if err != nil {
		return nil, err
	}

	if !data.VerifyTypeProtocolComunication(microservice.Type) {
		return nil, fmt.Errorf("Types of API protocols Does not exist")
	}

	microservices, err := usecase.microserviceRepository.InsertMicroservice(data)

	if err != nil {
		return nil, err
	}

	return microservices, nil
}

func (usecase *userUseCase) SearchPathService(path string, microservice *entities.Microservice) (*entities.Microservice, error) {
	data, err := usecase.microserviceRepository.FindPathMicroservice(path, microservice)

	if err != nil {
		return nil, fmt.Errorf("path is invalid")
	}

	return data, nil
}
